package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	"log"
)

func main() {
	pvk, err := crypto.GenerateKey() // we generate privateKey
	if err != nil {
		log.Fatal(err)
	}
	pData := crypto.FromECDSA(pvk)
	fmt.Println(hexutil.Encode(pData))

	pubData := crypto.FromECDSAPub(&pvk.PublicKey)
	fmt.Println(hexutil.Encode(pubData))

	fmt.Println(crypto.PubkeyToAddress(pvk.PublicKey).Hex())
}
