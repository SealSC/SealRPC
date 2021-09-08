package main

import (
	"SealRPC"
	"log"
)

func main() {
	ser := &SealRPC.EthService{}
	ser.EthGetTransactionReceiptFunc = func(TransactionHash string) (result SealRPC.EthGetTransactionReceiptResult, err error) {
		//Implementation you codes .....
		return
	}
	service := SealRPC.NewJsonRPCService(ser)
	err := service.Run(":8080")
	if err != nil {
		log.Fatalln("service run err:", err)
	}
}
