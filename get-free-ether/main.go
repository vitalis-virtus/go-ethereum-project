package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var url = "https://sepolia.infura.io/v3/9f66a25523c644fcaba087863b7a6058"

func main() {
	// we create two wallets
	// ks := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	// _, err := ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal("Error when create wallet")
	// }
	// _, err = ks.NewAccount("password")
	// if err != nil {
	// 	log.Fatal("Error when create wallet")
	// }

	// 93dd1c4800952758b6fa97a504fa4f77c09bbe4d
	// c757a52e98041a3a5af4f5b5eae60969b9f8e47f

	client, err := ethclient.DialContext(context.Background(), url)
	if err != nil {
		log.Fatal("Error to create a ether client: ", err)
	}
	defer client.Close()

	a1 := common.HexToAddress("93dd1c4800952758b6fa97a504fa4f77c09bbe4d")
	a2 := common.HexToAddress("c757a52e98041a3a5af4f5b5eae60969b9f8e47f")

	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal("Error when get the balance: ", err)
	}

	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal("Error when get the balance: ", err)
	}

	fmt.Println("Balance1: ", b1)
	fmt.Println("Balance2: ", b2)

}
