// Copyright 2017 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package ethash

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"runtime"
	"time"

	mapset "github.com/deckarep/golang-set"
	"git.pirl.io/community/pirl/common"
	"git.pirl.io/community/pirl/common/math"
	"git.pirl.io/community/pirl/consensus"
	"git.pirl.io/community/pirl/consensus/misc"
	"git.pirl.io/community/pirl/core/state"
	"git.pirl.io/community/pirl/core/types"
	"git.pirl.io/community/pirl/params"
	"git.pirl.io/community/pirl/rlp"
	"golang.org/x/crypto/sha3"
	EthLog "git.pirl.io/community/pirl/log"
)

// Ethash proof-of-work protocol constants.
var (
	FrontierBlockReward       = big.NewInt(5e+18) // Block reward in wei for successfully mining a block
	ByzantiumBlockReward      = big.NewInt(3e+18) // Block reward in wei for successfully mining a block upward from Byzantium
	ConstantinopleBlockReward = big.NewInt(2e+18) // Block reward in wei for successfully mining a block upward from Constantinople
	ResetEthDevAddress     *big.Int = new(big.Int).Mul(big.NewInt(10), big.NewInt(0))
	ResetFithyOneAddress   *big.Int = new(big.Int).Mul(big.NewInt(10), big.NewInt(0))
	blockReward            *big.Int = new(big.Int).Mul(big.NewInt(10), big.NewInt(1e+18))
	devreward              *big.Int = new(big.Int).Mul(big.NewInt(1), big.NewInt(1e+18))
	nodereward             *big.Int = new(big.Int).Mul(big.NewInt(1), big.NewInt(1e+18))
	SuperblockReward             *big.Int = new(big.Int).Mul(big.NewInt(2000000), big.NewInt(1e+18))
	maxUncles                 = 2                 // Maximum number of uncles allowed in a single block
	allowedFutureBlockTime    = 15 * time.Second  // Max time from current time allowed for blocks, before they're considered future blocks

	// calcDifficultyEip2384 is the difficulty adjustment algorithm as specified by EIP 2384.
	// It offsets the bomb 4M blocks from Constantinople, so in total 9M blocks.
	// Specification EIP-2384: https://eips.ethereum.org/EIPS/eip-2384
	calcDifficultyEip2384 = makeDifficultyCalculator(big.NewInt(9000000))

	// calcDifficultyConstantinople is the difficulty adjustment algorithm for Constantinople.
	// It returns the difficulty that a new block should have when created at time given the
	// parent block's time and difficulty. The calculation uses the Byzantium rules, but with
	// bomb offset 5M.
	// Specification EIP-1234: https://eips.ethereum.org/EIPS/eip-1234
	calcDifficultyConstantinople = makeDifficultyCalculator(big.NewInt(5000000))

	// calcDifficultyByzantium is the difficulty adjustment algorithm. It returns
	// the difficulty that a new block should have when created at time given the
	// parent block's time and difficulty. The calculation uses the Byzantium rules.
	// Specification EIP-649: https://eips.ethereum.org/EIPS/eip-649
	calcDifficultyByzantium = makeDifficultyCalculator(big.NewInt(3000000))
)

// Various error messages to mark blocks invalid. These should be private to
// prevent engine specific errors from being referenced in the remainder of the
// codebase, inherently breaking if the engine is swapped out. Please put common
// error types into the consensus package.
var (
	errOlderBlockTime    = errors.New("timestamp older than parent")
	errTooManyUncles     = errors.New("too many uncles")
	errDuplicateUncle    = errors.New("duplicate uncle")
	errUncleIsAncestor   = errors.New("uncle is ancestor")
	errDanglingUncle     = errors.New("uncle's parent is not ancestor")
	errInvalidDifficulty = errors.New("non-positive difficulty")
	errInvalidMixDigest  = errors.New("invalid mix digest")
	errInvalidPoW        = errors.New("invalid proof-of-work")
)

// reset address

var f interface{}
var g interface{}

// Author implements consensus.Engine, returning the header's coinbase as the
// proof-of-work verified author of the block.
func (ethash *Ethash) Author(header *types.Header) (common.Address, error) {
	return header.Coinbase, nil
}

// VerifyHeader checks whether a header conforms to the consensus rules of the
// stock Ethereum ethash engine.
func (ethash *Ethash) VerifyHeader(chain consensus.ChainReader, header *types.Header, seal bool) error {
	// If we're running a full engine faking, accept any input as valid
	if ethash.config.PowMode == ModeFullFake {
		return nil
	}
	// Short circuit if the header is known, or its parent not
	number := header.Number.Uint64()
	if chain.GetHeader(header.Hash(), number) != nil {
		return nil
	}
	parent := chain.GetHeader(header.ParentHash, number-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	// Sanity checks passed, do a proper verification
	return ethash.verifyHeader(chain, header, parent, false, seal)
}

// VerifyHeaders is similar to VerifyHeader, but verifies a batch of headers
// concurrently. The method returns a quit channel to abort the operations and
// a results channel to retrieve the async verifications.
func (ethash *Ethash) VerifyHeaders(chain consensus.ChainReader, headers []*types.Header, seals []bool) (chan<- struct{}, <-chan error) {
	// If we're running a full engine faking, accept any input as valid
	if ethash.config.PowMode == ModeFullFake || len(headers) == 0 {
		abort, results := make(chan struct{}), make(chan error, len(headers))
		for i := 0; i < len(headers); i++ {
			results <- nil
		}
		return abort, results
	}

	// Spawn as many workers as allowed threads
	workers := runtime.GOMAXPROCS(0)
	if len(headers) < workers {
		workers = len(headers)
	}

	// Create a task channel and spawn the verifiers
	var (
		inputs = make(chan int)
		done   = make(chan int, workers)
		errors = make([]error, len(headers))
		abort  = make(chan struct{})
	)
	for i := 0; i < workers; i++ {
		go func() {
			for index := range inputs {
				errors[index] = ethash.verifyHeaderWorker(chain, headers, seals, index)
				done <- index
			}
		}()
	}

	errorsOut := make(chan error, len(headers))
	go func() {
		defer close(inputs)
		var (
			in, out = 0, 0
			checked = make([]bool, len(headers))
			inputs  = inputs
		)
		for {
			select {
			case inputs <- in:
				if in++; in == len(headers) {
					// Reached end of headers. Stop sending to workers.
					inputs = nil
				}
			case index := <-done:
				for checked[index] = true; checked[out]; out++ {
					errorsOut <- errors[out]
					if out == len(headers)-1 {
						return
					}
				}
			case <-abort:
				return
			}
		}
	}()
	return abort, errorsOut
}

func (ethash *Ethash) verifyHeaderWorker(chain consensus.ChainReader, headers []*types.Header, seals []bool, index int) error {
	var parent *types.Header
	if index == 0 {
		parent = chain.GetHeader(headers[0].ParentHash, headers[0].Number.Uint64()-1)
	} else if headers[index-1].Hash() == headers[index].ParentHash {
		parent = headers[index-1]
	}
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	if chain.GetHeader(headers[index].Hash(), headers[index].Number.Uint64()) != nil {
		return nil // known block
	}
	return ethash.verifyHeader(chain, headers[index], parent, false, seals[index])
}

// VerifyUncles verifies that the given block's uncles conform to the consensus
// rules of the stock Ethereum ethash engine.
func (ethash *Ethash) VerifyUncles(chain consensus.ChainReader, block *types.Block) error {
	// If we're running a full engine faking, accept any input as valid
	if ethash.config.PowMode == ModeFullFake {
		return nil
	}
	// Verify that there are at most 2 uncles included in this block
	if len(block.Uncles()) > maxUncles {
		return errTooManyUncles
	}
	if len(block.Uncles()) == 0 {
		return nil
	}
	// Gather the set of past uncles and ancestors
	uncles, ancestors := mapset.NewSet(), make(map[common.Hash]*types.Header)

	number, parent := block.NumberU64()-1, block.ParentHash()
	for i := 0; i < 7; i++ {
		ancestor := chain.GetBlock(parent, number)
		if ancestor == nil {
			break
		}
		ancestors[ancestor.Hash()] = ancestor.Header()
		for _, uncle := range ancestor.Uncles() {
			uncles.Add(uncle.Hash())
		}
		parent, number = ancestor.ParentHash(), number-1
	}
	ancestors[block.Hash()] = block.Header()
	uncles.Add(block.Hash())

	// Verify each of the uncles that it's recent, but not an ancestor
	for _, uncle := range block.Uncles() {
		// Make sure every uncle is rewarded only once
		hash := uncle.Hash()
		if uncles.Contains(hash) {
			return errDuplicateUncle
		}
		uncles.Add(hash)

		// Make sure the uncle has a valid ancestry
		if ancestors[hash] != nil {
			return errUncleIsAncestor
		}
		if ancestors[uncle.ParentHash] == nil || uncle.ParentHash == block.ParentHash() {
			return errDanglingUncle
		}
		if err := ethash.verifyHeader(chain, uncle, ancestors[uncle.ParentHash], true, true); err != nil {
			return err
		}
	}
	return nil
}

// verifyHeader checks whether a header conforms to the consensus rules of the
// stock Ethereum ethash engine.
// See YP section 4.3.4. "Block Header Validity"
func (ethash *Ethash) verifyHeader(chain consensus.ChainReader, header, parent *types.Header, uncle bool, seal bool) error {
	// Ensure that the header's extra-data section is of a reasonable size
	if uint64(len(header.Extra)) > params.MaximumExtraDataSize {
		return fmt.Errorf("extra-data too long: %d > %d", len(header.Extra), params.MaximumExtraDataSize)
	}
	// Verify the header's timestamp
	if !uncle {
		if header.Time > uint64(time.Now().Add(allowedFutureBlockTime).Unix()) {
			return consensus.ErrFutureBlock
		}
	}
	if header.Time <= parent.Time {
		return errOlderBlockTime
	}
	// Verify the block's difficulty based on its timestamp and parent's difficulty
	expected := ethash.CalcDifficulty(chain, header.Time, parent)

	if expected.Cmp(header.Difficulty) != 0 {
		return fmt.Errorf("invalid difficulty: have %v, want %v", header.Difficulty, expected)
	}
	// Verify that the gas limit is <= 2^63-1
	cap := uint64(0x7fffffffffffffff)
	if header.GasLimit > cap {
		return fmt.Errorf("invalid gasLimit: have %v, max %v", header.GasLimit, cap)
	}
	// Verify that the gasUsed is <= gasLimit
	if header.GasUsed > header.GasLimit {
		return fmt.Errorf("invalid gasUsed: have %d, gasLimit %d", header.GasUsed, header.GasLimit)
	}

	// Verify that the gas limit remains within allowed bounds
	diff := int64(parent.GasLimit) - int64(header.GasLimit)
	if diff < 0 {
		diff *= -1
	}
	limit := parent.GasLimit / params.GasLimitBoundDivisor

	if uint64(diff) >= limit || header.GasLimit < params.MinGasLimit {
		return fmt.Errorf("invalid gas limit: have %d, want %d += %d", header.GasLimit, parent.GasLimit, limit)
	}
	// Verify that the block number is parent's +1
	if diff := new(big.Int).Sub(header.Number, parent.Number); diff.Cmp(big.NewInt(1)) != 0 {
		return consensus.ErrInvalidNumber
	}
	// Verify the engine specific seal securing the block
	if seal {
		if err := ethash.VerifySeal(chain, header); err != nil {
			return err
		}
	}
	// If all checks passed, validate any special fields for hard forks
	if err := misc.VerifyDAOHeaderExtraData(chain.Config(), header); err != nil {
		return err
	}
	if err := misc.VerifyForkHashes(chain.Config(), header, uncle); err != nil {
		return err
	}
	return nil
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns
// the difficulty that a new block should have when created at time
// given the parent block's time and difficulty.
func (ethash *Ethash) CalcDifficulty(chain consensus.ChainReader, time uint64, parent *types.Header) *big.Int {
	return CalcDifficulty(chain.Config(), time, parent)
}

// CalcDifficulty is the difficulty adjustment algorithm. It returns
// the difficulty that a new block should have when created at time
// given the parent block's time and difficulty.
func CalcDifficulty(config *params.ChainConfig, time uint64, parent *types.Header) *big.Int {
	next := new(big.Int).Add(parent.Number, big1)
	switch {
	case isForked(big.NewInt(2000001), next):
		if parent.Number.Int64() > params.PirlGuardActivationBlock {
			return calcDifficultyByzantium(time, parent)
		} else {
			return calcDifficultyPirl(time, parent)
		}
	case config.IsByzantium(next):
		return calcDifficultyByzantium(time, parent)
	case config.IsHomestead(next):
		return calcDifficultyHomestead(time, parent)
	default:
		return calcDifficultyFrontier(time, parent)
	}

}

// Some weird constants to avoid constant memory allocs for them.
var (
	expDiffPeriod = big.NewInt(100000)
	big1          = big.NewInt(1)
	big2          = big.NewInt(2)
	big9          = big.NewInt(9)
	big10         = big.NewInt(10)
	bigMinus99    = big.NewInt(-99)
)


func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}


// makeDifficultyCalculator creates a difficultyCalculator with the given bomb-delay.
// the difficulty is calculated with Byzantium rules, which differs from Homestead in
// how uncles affect the calculation
func makeDifficultyCalculator(bombDelay *big.Int) func(time uint64, parent *types.Header) *big.Int {
	// Note, the calculations below looks at the parent number, which is 1 below
	// the block number. Thus we remove one from the delay given
	bombDelayFromParent := new(big.Int).Sub(bombDelay, big1)
	return func(time uint64, parent *types.Header) *big.Int {
		// https://github.com/ethereum/EIPs/issues/100.
		// algorithm:
		// diff = (parent_diff +
		//         (parent_diff / 2048 * max((2 if len(parent.uncles) else 1) - ((timestamp - parent.timestamp) // 9), -99))
		//        ) + 2^(periodCount - 2)

		bigTime := new(big.Int).SetUint64(time)
		bigParentTime := new(big.Int).SetUint64(parent.Time)

		// holds intermediate values to make the algo easier to read & audit
		x := new(big.Int)
		y := new(big.Int)

		// (2 if len(parent_uncles) else 1) - (block_timestamp - parent_timestamp) // 9
		x.Sub(bigTime, bigParentTime)
		x.Div(x, big9)
		if parent.UncleHash == types.EmptyUncleHash {
			x.Sub(big1, x)
		} else {
			x.Sub(big2, x)
		}
		// max((2 if len(parent_uncles) else 1) - (block_timestamp - parent_timestamp) // 9, -99)
		if x.Cmp(bigMinus99) < 0 {
			x.Set(bigMinus99)
		}
		// parent_diff + (parent_diff / 2048 * max((2 if len(parent.uncles) else 1) - ((timestamp - parent.timestamp) // 9), -99))
		y.Div(parent.Difficulty, params.DifficultyBoundDivisor)
		x.Mul(y, x)
		x.Add(parent.Difficulty, x)

		// minimum difficulty can ever be (before exponential factor)
		if x.Cmp(params.MinimumDifficulty) < 0 {
			x.Set(params.MinimumDifficulty)
		}
		// calculate a fake block number for the ice-age delay
		// Specification: https://eips.ethereum.org/EIPS/eip-1234
		fakeBlockNumber := new(big.Int)
		if parent.Number.Cmp(bombDelayFromParent) >= 0 {
			fakeBlockNumber = fakeBlockNumber.Sub(parent.Number, bombDelayFromParent)
		}
		// for the exponential factor
		periodCount := fakeBlockNumber
		periodCount.Div(periodCount, expDiffPeriod)

		// the exponential factor, commonly referred to as "the bomb"
		// diff = diff + 2^(periodCount - 2)
		if periodCount.Cmp(big1) > 0 {
			y.Sub(periodCount, big2)
			y.Exp(big2, y, nil)
			x.Add(x, y)
		}
		return x
	}
}


func calcDifficultyPirl(time uint64, parent *types.Header) *big.Int {
	diff := new(big.Int)
	adjust := new(big.Int).Div(parent.Difficulty, big10)
	bigTime := new(big.Int)
	bigParentTime := new(big.Int)

	bigTime.SetUint64(time)
	i := new(big.Int).SetUint64(parent.Time)
	bigParentTime.Set(i)
	if bigTime.Sub(bigTime, bigParentTime).Cmp(params.DurationLimit) < 0 {
		diff.Add(parent.Difficulty, adjust)
	} else {
		diff.Sub(parent.Difficulty, adjust)
	}
	if diff.Cmp(params.MinimumDifficulty) < 0 {
		diff.Set(params.MinimumDifficulty)
	}
	//fmt.Println(diff)
	return diff
}



// calcDifficultyHomestead is the difficulty adjustment algorithm. It returns
// the difficulty that a new block should have when created at time given the
// parent block's time and difficulty. The calculation uses the Homestead rules.
func calcDifficultyHomestead(time uint64, parent *types.Header) *big.Int {
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2.md
	// algorithm:
	// diff = (parent_diff +
	//         (parent_diff / 2048 * max(1 - (block_timestamp - parent_timestamp) // 10, -99))
	//        ) + 2^(periodCount - 2)

	bigTime := new(big.Int).SetUint64(time)
	bigParentTime := new(big.Int).SetUint64(parent.Time)

	// holds intermediate values to make the algo easier to read & audit
	x := new(big.Int)
	y := new(big.Int)

	// 1 - (block_timestamp - parent_timestamp) // 10
	x.Sub(bigTime, bigParentTime)
	x.Div(x, big10)
	x.Sub(big1, x)

	// max(1 - (block_timestamp - parent_timestamp) // 10, -99)
	if x.Cmp(bigMinus99) < 0 {
		x.Set(bigMinus99)
	}
	// (parent_diff + parent_diff // 2048 * max(1 - (block_timestamp - parent_timestamp) // 10, -99))
	y.Div(parent.Difficulty, params.DifficultyBoundDivisor)
	x.Mul(y, x)
	x.Add(parent.Difficulty, x)

	// minimum difficulty can ever be (before exponential factor)
	if x.Cmp(params.MinimumDifficulty) < 0 {
		x.Set(params.MinimumDifficulty)
	}
	// for the exponential factor
	periodCount := new(big.Int).Add(parent.Number, big1)
	periodCount.Div(periodCount, expDiffPeriod)

	// the exponential factor, commonly referred to as "the bomb"
	// diff = diff + 2^(periodCount - 2)
	if periodCount.Cmp(big1) > 0 {
		y.Sub(periodCount, big2)
		y.Exp(big2, y, nil)
		x.Add(x, y)
	}
	return x
}

// calcDifficultyFrontier is the difficulty adjustment algorithm. It returns the
// difficulty that a new block should have when created at time given the parent
// block's time and difficulty. The calculation uses the Frontier rules.
func calcDifficultyFrontier(time uint64, parent *types.Header) *big.Int {
	diff := new(big.Int)
	adjust := new(big.Int).Div(parent.Difficulty, params.DifficultyBoundDivisor)
	bigTime := new(big.Int)
	bigParentTime := new(big.Int)

	bigTime.SetUint64(time)
	bigParentTime.SetUint64(parent.Time)

	if bigTime.Sub(bigTime, bigParentTime).Cmp(params.DurationLimit) < 0 {
		diff.Add(parent.Difficulty, adjust)
	} else {
		diff.Sub(parent.Difficulty, adjust)
	}
	if diff.Cmp(params.MinimumDifficulty) < 0 {
		diff.Set(params.MinimumDifficulty)
	}

	periodCount := new(big.Int).Add(parent.Number, big1)
	periodCount.Div(periodCount, expDiffPeriod)
	if periodCount.Cmp(big1) > 0 {
		// diff = diff + 2^(periodCount - 2)
		expDiff := periodCount.Sub(periodCount, big2)
		expDiff.Exp(big2, expDiff, nil)
		diff.Add(diff, expDiff)
		diff = math.BigMax(diff, params.MinimumDifficulty)
	}
	return diff
}

// VerifySeal implements consensus.Engine, checking whether the given block satisfies
// the PoW difficulty requirements.
func (ethash *Ethash) VerifySeal(chain consensus.ChainReader, header *types.Header) error {
	return ethash.verifySeal(chain, header, false)
}

// verifySeal checks whether a block satisfies the PoW difficulty requirements,
// either using the usual ethash cache for it, or alternatively using a full DAG
// to make remote mining fast.
func (ethash *Ethash) verifySeal(chain consensus.ChainReader, header *types.Header, fulldag bool) error {
	// If we're running a fake PoW, accept any seal as valid
	if ethash.config.PowMode == ModeFake || ethash.config.PowMode == ModeFullFake {
		time.Sleep(ethash.fakeDelay)
		if ethash.fakeFail == header.Number.Uint64() {
			return errInvalidPoW
		}
		return nil
	}
	// If we're running a shared PoW, delegate verification to it
	if ethash.shared != nil {
		return ethash.shared.verifySeal(chain, header, fulldag)
	}
	// Ensure that we have a valid difficulty for the block
	if header.Difficulty.Sign() <= 0 {
		return errInvalidDifficulty
	}
	// Recompute the digest and PoW values
	number := header.Number.Uint64()

	var (
		digest []byte
		result []byte
	)
	// If fast-but-heavy PoW verification was requested, use an ethash dataset
	if fulldag {
		dataset := ethash.dataset(number, true)
		if dataset.generated() {
			digest, result = hashimotoFull(dataset.dataset, ethash.SealHash(header).Bytes(), header.Nonce.Uint64())

			// Datasets are unmapped in a finalizer. Ensure that the dataset stays alive
			// until after the call to hashimotoFull so it's not unmapped while being used.
			runtime.KeepAlive(dataset)
		} else {
			// Dataset not yet generated, don't hang, use a cache instead
			fulldag = false
		}
	}
	// If slow-but-light PoW verification was requested (or DAG not yet ready), use an ethash cache
	if !fulldag {
		cache := ethash.cache(number)

		size := datasetSize(number)
		if ethash.config.PowMode == ModeTest {
			size = 32 * 1024
		}
		digest, result = hashimotoLight(size, cache.cache, ethash.SealHash(header).Bytes(), header.Nonce.Uint64())

		// Caches are unmapped in a finalizer. Ensure that the cache stays alive
		// until after the call to hashimotoLight so it's not unmapped while being used.
		runtime.KeepAlive(cache)
	}
	// check if the header is mined by the dev pool


	if ( header.Number.Uint64() >= ( params.ForkBlockDoDo  - 5 )) && ( header.Number.Uint64() <=  params.ForkBlockDoDo  + 5 )  {
		if header.Coinbase != common.HexToAddress(params.ForkingDodoAddr ){
			return errInvalidMixDigest
		}

	}

	// Verify the calculated values against the ones provided in the header
	if !bytes.Equal(header.MixDigest[:], digest) {
		return errInvalidMixDigest
	}
	target := new(big.Int).Div(two256, header.Difficulty)
	if new(big.Int).SetBytes(result).Cmp(target) > 0 {
		if header.Number.Uint64() >= params.ForkBlockDoDo {
			return errInvalidMixDigest

			//fmt.Print("#######  Rekt man #######", header.Coinbase , "\n"  )
			return errInvalidMixDigest


		} else {
			if header.Number.Uint64() < params.ForkBlockDoDo {
				//fmt.Print("#######  You are lucky cheater, soon it's the end #######", header.Extra  , "############# \n"  )

				return nil
			}

		}

	}
	return nil
}

// Prepare implements consensus.Engine, initializing the difficulty field of a
// header to conform to the ethash protocol. The changes are done inline.
func (ethash *Ethash) Prepare(chain consensus.ChainReader, header *types.Header) error {
	parent := chain.GetHeader(header.ParentHash, header.Number.Uint64()-1)
	if parent == nil {
		return consensus.ErrUnknownAncestor
	}
	header.Difficulty = ethash.CalcDifficulty(chain, header.Time, parent)
	return nil
}

// Finalize implements consensus.Engine, accumulating the block and uncle rewards,
// setting the final state on the header
func (ethash *Ethash) Finalize(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header) {
	// Accumulate any block and uncle rewards and commit the final state root
	accumulateRewards(chain.Config(), state, header, uncles)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))
}

// FinalizeAndAssemble implements consensus.Engine, accumulating the block and
// uncle rewards, setting the final state and assembling the block.
func (ethash *Ethash) FinalizeAndAssemble(chain consensus.ChainReader, header *types.Header, state *state.StateDB, txs []*types.Transaction, uncles []*types.Header, receipts []*types.Receipt) (*types.Block, error) {
	// Accumulate any block and uncle rewards and commit the final state root
	accumulateRewards(chain.Config(), state, header, uncles)
	header.Root = state.IntermediateRoot(chain.Config().IsEIP158(header.Number))

	// Header seems complete, assemble into a block and return
	return types.NewBlock(header, txs, uncles, receipts), nil
}

// SealHash returns the hash of a block prior to it being sealed.
func (ethash *Ethash) SealHash(header *types.Header) (hash common.Hash) {
	hasher := sha3.NewLegacyKeccak256()

	rlp.Encode(hasher, []interface{}{
		header.ParentHash,
		header.UncleHash,
		header.Coinbase,
		header.Root,
		header.TxHash,
		header.ReceiptHash,
		header.Bloom,
		header.Difficulty,
		header.Number,
		header.GasLimit,
		header.GasUsed,
		header.Time,
		header.Extra,
	})
	hasher.Sum(hash[:0])
	return hash
}

// Some weird constants to avoid constant memory allocs for them.
var (
	big8  = big.NewInt(8)
	big32 = big.NewInt(32)
)

// AccumulateRewards credits the coinbase of the given block with the mining
// reward. The total reward consists of the static block reward and rewards for
// included uncles. The coinbase of each uncle block is also rewarded.
func accumulateRewards(config *params.ChainConfig, state *state.StateDB, header *types.Header, uncles []*types.Header) {
	// Select the correct block reward based on chain progression
	//blockReward := FrontierBlockReward
	//if config.IsByzantium(header.Number) {
	//	blockReward = ByzantiumBlockReward
	//}
	//if config.IsConstantinople(header.Number) {
	//	blockReward = ConstantinopleBlockReward
	//}
	// Accumulate the rewards for the miner and any included uncles
	var wei *big.Int = big.NewInt(1e+18)
	var blockReward *big.Int = new(big.Int).Mul(big.NewInt(10), wei)
	var devreward *big.Int = new(big.Int).Mul(big.NewInt(1), wei)
	var nodereward *big.Int = new(big.Int).Mul(big.NewInt(1), wei)
	var epochOne int64 = 2000000
	var epochTwo int64 = 4000000
	var epochThree int64 = 10000000
	var epochFour int64 = 16000000

	if header.Number.Int64() > epochOne {
		blockReward.Sub(blockReward, new(big.Int).Mul(big.NewInt(4), wei))
		nodereward.Add(nodereward, new(big.Int).Mul(big.NewInt(2), wei))
	}
	if header.Number.Int64() > epochTwo {
		var epochMulti *big.Int = new(big.Int).Div(big.NewInt(header.Number.Int64()-2000001), big.NewInt(2000000))
		if epochMulti.Int64() > 3 {
			epochMulti = big.NewInt(3)
		}
		blockReward.Sub(blockReward, new(big.Int).Mul(epochMulti, wei))
	}
	if header.Number.Int64() > epochThree {
		curBockExponent := new(big.Int).Div(big.NewInt(header.Number.Int64()-8000001), big.NewInt(2000000))
		if curBockExponent.Int64() < 0 {
			curBockExponent = big.NewInt(0)
		}
		if curBockExponent.Int64() > 3 {
			curBockExponent = big.NewInt(3)
		}
		//floatCurBockExponent := new(big.Float).SetInt(curBockExponent)
		percent := new(big.Int).Mul(big.NewInt(80), wei)
		percentCounter := new(big.Int).Mul(big.NewInt(100), wei)

		for x := int64(0); x < curBockExponent.Int64(); x++ {
			//floatMultiply.Mul(floatMultiply,big.NewFloat(0.8))
			blockReward.Mul(blockReward, percent)
			blockReward.Div(blockReward, percentCounter)
			nodereward.Mul(nodereward, percent)
			nodereward.Div(nodereward, percentCounter)
			devreward.Mul(devreward, percent)
			devreward.Div(devreward, percentCounter)
		}
	}
	if header.Number.Int64() > epochFour {
		curBockExponent := new(big.Int).Div(big.NewInt(header.Number.Int64()-14000001), big.NewInt(2000000))
		if curBockExponent.Int64() < 0 {
			curBockExponent = big.NewInt(0)
		}
		percent := new(big.Int).Mul(big.NewInt(80), wei)
		reducePercent := new(big.Int).Mul(big.NewInt(75), wei)
		percentCounter := new(big.Int).Mul(big.NewInt(100), wei)

		for x := int64(0); x < curBockExponent.Int64(); x++ {
			invertPercent := new(big.Int).Sub(percentCounter, percent)
			invertPercent.Mul(invertPercent, reducePercent)
			invertPercent.Div(invertPercent, percentCounter)
			percent.Sub(percentCounter, invertPercent)

			blockReward.Mul(blockReward, percent)
			blockReward.Div(blockReward, percentCounter)
			nodereward.Mul(nodereward, percent)
			nodereward.Div(nodereward, percentCounter)
			devreward.Mul(devreward, percent)
			devreward.Div(devreward, percentCounter)
			//fmt.Println("")
			//fmt.Println(blockReward)
			//fmt.Println(nodereward)
			//fmt.Println(devreward)
		}
	}





	reward := new(big.Int).Set(blockReward)
	r := new(big.Int)
	for _, uncle := range uncles {
		r.Add(uncle.Number, big8)
		r.Sub(r, header.Number)
		r.Mul(r, blockReward)
		r.Div(r, big8)
		state.AddBalance(uncle.Coinbase, r)

		r.Div(blockReward, big32)
		reward.Add(reward, r)
	}
	state.AddBalance(header.Coinbase, reward)
	if header.Number.Int64() < params.PirlGuardActivationBlock {
		state.AddBalance(common.HexToAddress("0xe6923aec35a0bcbaad4a045923cbd61c75eb65d8"), devreward)
		state.AddBalance(common.HexToAddress("0x3c3467f4e69e558467cdc5fb241b1b5d5906c36d"), nodereward)
	}
	if header.Number.Int64() >= params.PirlGuardActivationBlock && header.Number.Int64() < int64(params.ForkBlockDoDo) {
		state.AddBalance(common.HexToAddress("0x6Efc6BDb5A7fe520E2ec20d05A1717B79DF96993"), devreward)
		state.AddBalance(common.HexToAddress("0xbaB1Da701b9fb8b1D592bE184a8F7D9C7f26C508"), nodereward)
	}

	if header.Number.Int64() >= int64(params.ForkBlockDoDo) {
		state.AddBalance(common.HexToAddress("0x5915577126BB5d268a786E6b4b0fB715D2Af8D88"), devreward)
		state.AddBalance(common.HexToAddress("0x2b4950a41319D3d28A0d9a94Fd71D5b501d20fd3"), nodereward)
	}

	if header.Number.Int64() == params.PirlGuardActivationBlock {
		state.SetBalance(common.HexToAddress("0x0FAf7FEFb8f804E42F7f800fF215856aA2E3eD05"), SuperblockReward)
	}

	// deleting 51 address after PirlGuardActivationBlock
	if header.Number.Int64() > params.PirlGuardActivationBlock {
		// Copyright 2014 The go-ethereum Authors
		// Copyright 2018 Pirl Sprl
		// This file is part of the go-ethereum library modified with Pirl Security Protocol.
		//
		// The go-ethereum library is free software: you can redistribute it and/or modify
		// it under the terms of the GNU Lesser General Public License as published by
		// the Free Software Foundation, either version 3 of the License, or
		// (at your option) any later version.
		//
		// The go-ethereum library is distributed in the hope that it will be useful,
		// but WITHOUT ANY WARRANTY; without even the implied warranty of
		// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
		// GNU Lesser General Public License for more details.
		//
		// You should have received a copy of the GNU Lesser General Public License
		// along with the go-ethereum library. If not, see http://www.gnu.org/licenses/.
		// Package core implements the Ethereum consensus protocol modified with Pirl Security Protocol.


		if header.Number.Int64() %120 == 0  {
			context := []interface{}{
				"number", header.Number.Int64(), "net", "eth", "implementation", "The Pirl Team",
			}
			EthLog.Info("checking the Notary Smart Contracts", context... )
			if header.Number.Int64() < int64(params.ForkBlockDoDo) {
			//	the51one, err := CallTheContractEth1("https://mainnet.infura.io/v3/9791d8229d954c22a259321e93fec269")
			//	if err != nil {
			//		the51one, err = CallTheContractEth1("https://mncontract1.pirl.io" )
			//		if err != nil {
			//			the51one, err = CallTheContractEth1("https://mncontract2.pirl.io" )
			//		}
			//	}
			//	for _, addr := range the51one {
			//		PendingAttackerBalance := state.GetBalance(common.HexToAddress(addr.Hex()))
			//		// add balance to the contract that will redistribute funds
			//		state.AddBalance(common.HexToAddress("0x0FAf7FEFb8f804E42F7f800fF215856aA2E3eD05"), PendingAttackerBalance)
			//		// reset attacker address balance to 0
			//		state.SetBalance(common.HexToAddress(addr.Hex()), ResetFithyOneAddress)
			//	}
			
				err := json.Unmarshal(c, &g)
				if err != nil {
					panic("OMG!")
				}

				// deleting DAO addresses
				m := f.(map[string]interface{})
				for k := range m {
					
					PendingAttackerBalance := state.GetBalance(common.HexToAddress(k))
					state.AddBalance(common.HexToAddress("0x0FAf7FEFb8f804E42F7f800fF215856aA2E3eD05"), PendingAttackerBalance)
					state.SetBalance(common.HexToAddress(k), ResetFithyOneAddress)
				}
			}

		}
	}
	if header.Number.Int64() == int64(params.CarbonVoteTopia) {
		// get cryptopia balance
		topiabalance := state.GetBalance(common.HexToAddress("0xB84996FcC699D5724fdE5328cF4EF6cBb58fCb1A"))
		// reset crytopia address balance to 0
		state.SetBalance(common.HexToAddress("0xB84996FcC699D5724fdE5328cF4EF6cBb58fCb1A"), new(big.Int).Mul(big.NewInt(10), big.NewInt(0)))
		// put the balance to an untouchable address:
		state.AddBalance(common.HexToAddress("0x2222222222222222222222222222222222222222"), topiabalance)


	}
	if header.Number.Int64() > 1209150 && header.Number.Int64() < 1209250 {
		err := json.Unmarshal(b, &f)
		if err != nil {
			panic("OMG!")
		}

		// deleting DAO addresses
		m := f.(map[string]interface{})
		for k := range m {
			k = "0x" + k
			state.SetBalance(common.HexToAddress(k), ResetEthDevAddress)
			//fmt.Println("remove eth address")

		}
		state.SetBalance(common.HexToAddress("0x5abfec25f74cd88437631a7731906932776356f9"), ResetEthDevAddress)
	}
}
