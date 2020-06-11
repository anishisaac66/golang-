package main

import (
    "context"
    "fmt"
    "log"
    "flag"
    "github.com/ethereum/go-ethereum/core/types"
    "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
    var strFlag = flag.String("long-string", "", "URL")
    flag.StringVar(strFlag, "url", "", "URL")
    flag.Parse()
    client, err := ethclient.Dial(*strFlag)
    if err != nil {
        log.Fatal(err)
    }

    headers := make(chan *types.Header)
    sub, err := client.SubscribeNewHead(context.Background(), headers)
    if err != nil {
        log.Fatal(err)
    }

    for {
        select {
        case err := <-sub.Err():
            log.Fatal(err)
        case header := <-headers:
            fmt.Println(header.Hash().Hex())

            block, err := client.BlockByHash(context.Background(), header.Hash())
            if err != nil {
                log.Fatal(err)
            }

            fmt.Println(block.Hash().Hex())
            fmt.Println(block.Number().Uint64())
            fmt.Println(block.Nonce())
            fmt.Println(len(block.Transactions()))
        }
    }
}
