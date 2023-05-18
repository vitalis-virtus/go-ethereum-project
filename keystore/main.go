package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// generate file with address
	// key := keystore.NewKeyStore("./wallets", keystore.StandardScryptN, keystore.StandardScryptP)
	var password = "password"
	// addr, err := key.NewAccount(password)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(addr.Address)

	// get private, public keys and address from file
	b, err := ioutil.ReadFile("./wallets/UTC--2023-05-18T11-02-45.530355800Z--2e597af042ad5bf618f98ec99791a10f754b8c32")
	if err != nil {
		log.Fatal("Error read from file: ", err)
	}
	key, err := keystore.DecryptKey(b, password)
	if err != nil {
		log.Fatal("Error to decrypt file: ", err)
	}
	privData := crypto.FromECDSA(key.PrivateKey)
	fmt.Println("Priv: ", hexutil.Encode(privData))

	pubData := crypto.FromECDSAPub(&key.PrivateKey.PublicKey)
	fmt.Println("Pub: ", hexutil.Encode(pubData))

	addr := crypto.PubkeyToAddress(key.PrivateKey.PublicKey).Hex()
	fmt.Println("Addr ", addr)
}
