# SealRPC


## About
SealRPC service and client

## Files

```
├── README.md  
├── client.go  #SDK client
├── common.go  #SDK interface and structs
├── examples   #Some demos
├── go.mod    
├── go.sum
├── jsonrpc.go  #jsonRPC API using net/http
├── service.go  #Register and service these functions 
├── tools.go    #Tools library
```

## Install

```
 go get github.com/SealSC/SealRPC
```

## Use

### server
```go
//1. new EthService{}
ser := &SealRPC.EthService{}
//2. Register service function
ser.EthGetTransactionReceiptFunc = func (TransactionHash string) (result SealRPC.EthGetTransactionReceiptResult, err error) {
//Implementation you codes .....
return
}
//3. Binding external access protocol(Currently only ` jsonRPC`)
rpcSer := SealRPC.NewJsonRPCService(ser)
//4. Call 'rpcSer.run()' to start your service
err := rpcSer.Run(":8080")
if err != nil {
log.Fatalln("service run err:", err)
}

```

### client

```go
//1. Creat client using url
client := SealRPC.NewEthRPCClient("http://127.0.0.1:8080")
//2. Use the created 'client' functions
number, err := client.EthGetBlockByNumber("0x1", true)
fmt.Println(number, err)
index, err := client.EthGetBlockTransactionCountByHash("0xe9dc52ef255d3b8ee11b35acc4b25f41ce096c71392aafe58907df9e1dbaaa4b")
fmt.Println(index, err)
```

##Routing rules for jsonRPC functions

1. Read jsonRPC `RPCRequest.method` 
2. Call `tool.toUp(methodName)` to perform (camel-case) operation on methodName
3. Call `reflect.ValueOf().MethodByName` to get the reflect.method variable
4. Assemble params and call the function `reflect.Method.Call()` to return returns
