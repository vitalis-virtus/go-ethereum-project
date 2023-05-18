package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var url = "https://sepolia.infura.io/v3/9f66a25523c644fcaba087863b7a6058"

func main() {
	client, err := ethclient.DialContext(context.Background(), url)
	if err != nil {
		log.Fatal("Error when create client: ", err)
	}
	defer client.Close()

	a1 := common.HexToAddress("64d038f1de9708613a03ce87e145afe178188358")
	a2 := common.HexToAddress("5871b84720e30bee455929b4c86b08a50a06ea10")

	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal("Error when get client balance: ", err)
	}
	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal("Error when get client balance: ", err)
	}

	fmt.Println("Balance1: ", b1)
	fmt.Println("Balance2: ", b2)

	nonce, err := client.PendingNonceAt(context.Background(), a1)
	if err != nil {
		log.Fatal("Error when get the nonce")
	}
	amount := big.NewInt(1000000000000)
	price, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Error when suggest gas price: ", err)
	}
	tx := types.NewTransaction(nonce, a2, amount, 21000, price, nil)

	// we need t sign this tx by private key
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("Error when get chain ID: ", err)
	}

	b, err := ioutil.ReadFile("./wallets/UTC--2023-05-18T13-19-09.852416200Z--64d038f1de9708613a03ce87e145afe178188358")
	if err != nil {
		log.Fatal("Error read from file: ", err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal("Error to decrypt file: ", err)
	}

	fmt.Println("key: ", key.PrivateKey)

	types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)

	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal("Error when make a transaction: ", err)
	}
	fmt.Println("tx sent: ", tx.Hash().Hex())
}
