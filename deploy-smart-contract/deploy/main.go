package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"

	todo "go-ethereum-project/deploy-smart-contract/gen"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	b, err := ioutil.ReadFile("../wallets/UTC--2023-05-18T15-17-12.927298000Z--a6e74cb9689c6ed9f958835a92b89782ee994fdd")
	if err != nil {
		log.Fatal("Error to open file with wallet: ", err)
	}

	key, err := keystore.DecryptKey(b, "password")
	if err != nil {
		log.Fatal("Error to decrypt key: ", err)
	}

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/9f66a25523c644fcaba087863b7a6058")
	if err != nil {
		log.Fatal("Error to create clint: ", err)
	}
	defer client.Close()

	// get the address of keystore wallet
	address := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)

	// we request nonce from wallet because we dont know how many transactions we done
	nonce, err := client.PendingNonceAt(context.Background(), address)
	if err != nil {
		log.Fatal("Error when get the nonce: ", err)
	}

	// we suggest the gasPrice what we need to deploy our contract
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Error when get the gas price: ", err)
	}

	// chainID is id of the network
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("Error when get chain ID: ", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal("Error when create a transaction signer: ", err)
	}
	auth.GasPrice = gasPrice
	auth.GasLimit = uint64(3000000)
	auth.Nonce = big.NewInt(int64(nonce))
	a, tx, _, err := todo.DeployTodo(auth, client)
	if err != nil {
		log.Fatal("Error when deploy a new Ethereum contract: ", err)
	}
	fmt.Println("Address of contract: ", a.Hex())
	fmt.Println("Transaction: ", tx.Hash().Hex())
}
