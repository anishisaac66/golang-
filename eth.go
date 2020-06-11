
package main

import(
    "fmt"
    "log"

    "github.com/ethereum/go-ethereum/rpc"
)

type Block struct {
    Number string
}

func main() {
  // Connect the client
  client, err := rpc.Dial("http://52.44.214.192:22000")
  if err != nil {
    log.Fatalf("could not create ipc client: %v", err)
  }

  var lastBlock Block
  err = client.Call(&lastBlock, "eth_getBlockByNumber", "latest", true)
  if err != nil {
      fmt.Println("can't get latest block:", err)
      return
  }

  // Print events from the subscription as they arrive.
  fmt.Printf("latest block: %v\n", lastBlock.Number)
}
