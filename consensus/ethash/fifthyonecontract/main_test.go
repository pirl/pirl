//package fifthyonecontracttest
package main_test

import (
	Black "./redlist"
	"fmt"
	"git.pirl.io/community/pirl/accounts/abi/bind"
	"git.pirl.io/community/pirl/common"
	"git.pirl.io/community/pirl/ethclient"
	"log"
	"os"

	//"os"
)
// eth contract 0xdc427e8c5390e05cc4dd9f35ffd3b5c855a7ac26
// pirl contract 0x936C02bAf9A9A2efFC3deFDB8eAdcc3bFEeA8Ef0
func CallTheContractEth1(contractendpoint string) ([]common.Address, error) {

	endPoint := contractendpoint
	conn, err := ethclient.Dial(endPoint)
	//log.Printf("Connected to the Eth client")
	if err != nil {
		log.Printf("Failed to connect to the Eth client: %v", err)
		return nil, err
	}
	// Address Blacklist contract
	contract, err := Black.NewRedlistCaller(common.HexToAddress("0xdc427e8c5390e05cc4dd9f35ffd3b5c855a7ac26"), conn) // for dev
	//contract, err := Black.NewRedlistCaller(common.HexToAddress("0x5A5ddC83432DEf31F674Af38E5D0D02445d8Fc03"), conn)
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

func main() {
	fmt.Print(CallTheContractPirl())
	mycontract, err := CallTheContractEth1("https://mainnet.infura.io/v3/9791d8229d954c22a259321e93fec269")
	if err != nil {
		mycontract, _ = CallTheContractEth1("https://mainnet.infura.io/v3/9791d8229d954c22a259321e93fec269")
	}

	log.Println(mycontract)

}
