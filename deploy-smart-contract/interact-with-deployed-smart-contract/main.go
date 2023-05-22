package main

import (
	"context"
	"fmt"
	todo "go-ethereum-project/deploy-smart-contract/gen"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
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

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal("Error when get the gas price: ", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal("Error when get chain ID: ", err)
	}

	cAdd := common.HexToAddress("address of the contract")

	t, err := todo.NewTodo(cAdd, client)
	if err != nil {
		log.Fatal("Error to create todo structure: ", err)
	}

	tx, err := bind.NewKeyedTransactorWithChainID(key.PrivateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	tx.GasLimit = 3000000
	tx.GasPrice = gasPrice

	// create taks
	trans, err := t.Add(tx, "First task")
	if err != nil {
		log.Fatal("Error to add transaction: ", err)
	}

	// show hash of created task
	fmt.Println(trans.Hash())

	// list all tasks
	add := crypto.PubkeyToAddress(key.PrivateKey.PublicKey)
	tasks, err := t.List(&bind.CallOpts{
		From: add,
	})
	if err != nil {
		log.Fatal("Error to list tasks: ", err)
	}

	// show list of all tasks
	fmt.Println(tasks)

	// update task
	tra, err := t.Update(tx, big.NewInt(0), "update task content")
	if err != nil {
		log.Fatal("Error to update task: ", err)
	}

	// show hash of updated task
	fmt.Println("Toggle tx: ", tra.Hash())

	// toggle task
	trToggle, err := t.Toggle(tx, big.NewInt(0))
	if err != nil {
		log.Fatal("Error to toggle task: ", err)
	}

	// show hash of toggled task
	fmt.Println("Toggled tx: ", trToggle.Hash())

	// delete the task
	trDel, err := t.Remove(tx, big.NewInt(0))
	if err != nil {
		log.Fatal("Error to delete the task: ", err)
	}

	// show hash of deleted task
	fmt.Println("Toggled tx: ", trDel.Hash())
}
