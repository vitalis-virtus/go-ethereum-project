package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// we will interact with mainnet
var infuraURL = "https://mainnet.infura.io/v3/9f66a25523c644fcaba087863b7a6058"
var ganacheURL = "http://localhost:8545"

func main() {
	client, err := ethclient.DialContext(context.Background(), infuraURL) // we creating a client
	if err != nil {
		log.Fatalf("Error to create a ether client: %v", err)
	}
	defer client.Close()
	block, err := client.BlockByNumber(context.Background(), nil) // if we pass nill we will get the last block mined	 on Ethereum blockchain
	if err != nil {
		log.Fatalf("Error to get a block: %v", err)
	}
	fmt.Println("The block number: ", block.Number())

	addr := "0xfef7ddaae3e406c3fd5fa0a731c5adbae7afd7fa"
	address := common.HexToAddress(addr)

	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatalf("Error get the balance: %v", err)
	}

	fBalance := new(big.Float)
	fBalance.SetString(balance.String())
	fmt.Println("The balance: ", balance)
	fmt.Println(fBalance)
	balanceEth := new(big.Float).Quo(fBalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(balanceEth)
}
