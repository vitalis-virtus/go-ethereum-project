package main

import (
	"context"
	"fmt"
	"log"

	// "github.com/ethereum/go-ethereum/accounts/keystore"
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

	client, err := ethclient.DialContext(context.Background(), url)
	if err != nil {
		log.Fatal("Error to create a ether client: ", err)
	}
	defer client.Close()

	a1 := common.HexToAddress("a6e74cb9689c6ed9f958835a92b89782ee994fdd")
	a2 := common.HexToAddress("ee5be175a72a23b08ef163b61e1b95a76c616008")

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
