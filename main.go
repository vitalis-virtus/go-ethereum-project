package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

// we will interact with mainnet
var infuraURL = "https://mainnet.infura.io/v3/9f66a25523c644fcaba087863b7a6058"
var ganacheURL = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), ganacheURL) // we creating a client
	if err != nil {
		log.Fatal("Error to create a ether client: %v", err)
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil) // if we pass nill we will get the last block mined	 on Ethereum blockchain
	if err != nil {
		log.Fatal("Error to get a block: %v", err)
	}
	fmt.Println(block.Number())
}
