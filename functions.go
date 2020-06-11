package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
        "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://52.44.214.192:22000")
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := big.NewInt(346)

	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	for _, tx := range block.Transactions() {
		fmt.Printf("TX Hash: %s\n", tx.Hash().Hex())
		fmt.Printf("TX Value: %s\n", tx.Value().String())
		fmt.Printf("TX Gas: %d\n", tx.Gas())
		fmt.Printf("TX Gas Price: %d\n", tx.GasPrice().Uint64())
		fmt.Printf("TX Nonce: %d\n", tx.Nonce())
		fmt.Printf("TX Data: %v\n", tx.Data())
		fmt.Printf("TX To: %s\n", tx.To().Hex())

		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Receipt Status: %d\n", receipt.Status)
		fmt.Println("---")
	}
}
