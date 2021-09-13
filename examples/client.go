package main

import (
	"SealRPC"
	"fmt"
)

func main() {
	client := SealRPC.NewJsonRPCClient("http://127.0.0.1:8080")
	number, err := client.EthGetBlockByNumber("0x1", true)
	fmt.Println(number, err)
	index, err := client.EthGetBlockTransactionCountByHash("0xe9dc52ef255d3b8ee11b35acc4b25f41ce096c71392aafe58907df9e1dbaaa4b")
	fmt.Println(index, err)
}
