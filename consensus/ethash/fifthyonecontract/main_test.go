//package fifthyonecontract
package main

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
func CallTheContractEth() []common.Address{
//func main() {
	//endPoint := os.Getenv("HOME") + "/.pirl/pirl.ipc"
	endPoint := "https://mainnet.infura.io/v3/9791d8229d954c22a259321e93fec269"

	//home_pirl := os.Getenv("HOME")
	//endPoint :=  home_pirl + "/Library/Pirl/pirl.ipc"
	conn, err := ethclient.Dial(endPoint)

	if err != nil {
		log.Fatalf("Failed to connect to the Pirl client: %v", err)
	}
	/// Address Blacklist contract
	contract, err := Black.NewRedListCaller(common.HexToAddress("0xdc427e8c5390e05cc4dd9f35ffd3b5c855a7ac26"), conn)
	if err != nil {
		log.Fatalf("Failed to instantiate a trigger contract: %v", err)
	}

	session := Black.RedListCallerSession{
		Contract: contract,
		CallOpts: bind.CallOpts{
			Pending: true,
		},
	}
		isBanned, err := session.GetAllAddresses()
		if err != nil {
		log.Fatalf("Failed to connect to the Pirl client: %v", err)
	}
	return isBanned
	}

func CallTheContractPirl() ([]common.Address, error ){
	//func main() {
	//endPoint := os.Getenv("HOME") + "/.pirl/pirl.ipc"
	//endPoint := "https://wallrpc.pirl.io"

	home_pirl := os.Getenv("HOME")
	endPoint :=  home_pirl + "/Library/Pirl/pirl.ip"
	conn, err := ethclient.Dial(endPoint)

	if err != nil {
		log.Printf("Failed to connect to the Pirl client")

		return  nil , err
	}
	/// Address Blacklist contract
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
		log.Fatalf("Failed to connect to the Pirl client: %v", err)
		CallTheContractEth()
	}

	return isBanned, err
}
func main() {
	//fmt.Print(CallTheContractPirl())
	mycontract, err := CallTheContractPirl()
	if err != nil {
		mycontract = CallTheContractEth()
	}

	for _, addr := range mycontract {

		fmt.Println(addr.Hex())

	}
}
