package ethash

import (
	Black "git.pirl.io/community/pirl/consensus/ethash/redlist"
	"git.pirl.io/community/pirl/accounts/abi/bind"
	"git.pirl.io/community/pirl/common"
	"git.pirl.io/community/pirl/ethclient"
	"log"
	"os"
)

// contact smart contract eth
func CallTheContractEth() ([]common.Address, error) {

	endPoint := "https://mainnet.infura.io/v3/9791d8229d954c22a259321e93fec269"
	conn, err := ethclient.Dial(endPoint)
	//log.Printf("Connected to the Eth client")
	if err != nil {
		log.Printf("Failed to connect to the Eth client: %v", err)
		return nil, err
	}
	// Address Blacklist contract
	contract, err := Black.NewRedListCaller(common.HexToAddress("0xdc427e8c5390e05cc4dd9f35ffd3b5c855a7ac26"), conn)
	if err != nil {
		log.Printf("Failed to instantiate a trigger contract: %v", err)
	}

	session := Black.RedListCallerSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	isBanned, err := session.GetAllAddresses()
	if err != nil {
		log.Printf("Failed to connect to the Pirl client: %v", err)
	}
	return isBanned, err


}

// contact smart contract Pirl
func CallTheContractPirl() ([]common.Address, error) {

	endPoint := os.Getenv("HOME") + "/.pirl/pirl.ipc"
	conn, err := ethclient.Dial(endPoint)

	if err != nil {
		log.Printf("Failed to connect to the pirl client: %v", err)
		return nil, err
	}
	// Address Blacklist contract
	contract, err := Black.NewRedListCaller(common.HexToAddress("0x936C02bAf9A9A2efFC3deFDB8eAdcc3bFEeA8Ef0"), conn)
	if err != nil {
		log.Printf("Failed to instantiate a trigger contract: %v", err)
	}

	session := Black.RedListCallerSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
	isBanned, err := session.GetAllAddresses()
	if err != nil {
		log.Printf("Failed to connect to the Pirl client: %v", err)
	}

	return isBanned, err
}