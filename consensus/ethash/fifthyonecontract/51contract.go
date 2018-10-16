package fifthyonecontract


import (
	"github.com/ethereum/go-ethereum/ethclient"
	"os"
)


func CallTheContract() {
	endPoint := os.Getenv("HOME") + "/.pirl/pirl.ipc"
	client, err = ethclient.Dial(endPoint)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}


}