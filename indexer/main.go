package main

import (
	"context"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
)

func getEnv() string {
	rpcUrl := os.Getenv("RPC_URL")
	if len(rpcUrl) <= 0 {
		log.Fatal("No RPC_URL set or could not be found, please set in the .env file and try again")
	}

	return rpcUrl
}

func main() {
	rpcUrl := getEnv()

	client, err := ethclient.Dial(rpcUrl)

	if err != nil {
		log.Fatal("Error when attempting to call out through etherium client, with error", err.Error())
	}

	// If for some reason we crash or the app closes forcefully then we gracefully close the client.
	defer client.Close()

	blockNum, err := client.BlockNumber(context.Background())

	log.Printf("Block Number: %d", blockNum)
}
