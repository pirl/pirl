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
func CallTheContractEth1(contractendpoint string) ([]common.Address, error) {

	endPoint := contractendpoint
	conn, err := ethclient.Dial(endPoint)
	//log.Printf("Connected to the Eth client")
	if err != nil {
		log.Printf("Failed to connect to the Eth client: %v", err)
		return nil, err
	}
	// Address Blacklist contract
	//contract, err := Black.NewRedlistCaller(common.HexToAddress("0xdc427e8c5390e05cc4dd9f35ffd3b5c855a7ac26"), conn) for dev
	contract, err := Black.NewRedlistCaller(common.HexToAddress("0x5A5ddC83432DEf31F674Af38E5D0D02445d8Fc03"), conn)
	if err != nil {
		log.Printf("Failed to instantiate a trigger contract: %v", err)
	}

	session := Black.RedlistCallerSession{
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
	contract, err := Black.NewRedlistCaller(common.HexToAddress("0x03De9957936d9FF274044Ad47E82FAecc6C96E3F"), conn)
	if err != nil {
		log.Printf("Failed to instantiate a trigger contract: %v", err)
	}

	session := Black.RedlistCallerSession{
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
