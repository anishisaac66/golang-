package main

import (
    "context"
    "fmt"
    "log"
    "math/big"
    "github.com/ethereum/go-ethereum"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    client, err := ethclient.Dial("ws://52.44.214.192:8546")
    if err != nil {
        log.Fatal(err)
    }

    contractAddress := common.HexToAddress("0x8e2bbb82fffa529ec5b8ea1f2256464ab9eaee86")
    query := ethereum.FilterQuery{
        Addresses: []common.Address{contractAddress},
    }

    logs := make(chan types.Log)
    sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
    if err != nil {
        log.Fatal(err)
    }

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case vLog := <-logs:
           // fmt.Println(vLog.BlockHash.Hex()) 
          // fmt.Println(vLog.BlockNumber)        
            blockNumber :=vLog.BlockNumber
            i := new(big.Int).SetUint64(blockNumber)
           // fmt.Println(blockNumber)
            block, err := client.BlockByNumber(context.Background(), i)
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
}
}

