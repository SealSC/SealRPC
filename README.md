# SealRPC

## About

SealRPC service and client

- [ ] jsonRPC1.0 supported
- [x] jsonRPC2.0 supported

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

## Routing rules for jsonRPC functions

1. Read jsonRPC `RPCRequest.method`
2. Call `tool.toUp(methodName)` to perform (camel-case) operation on methodName
3. Call `reflect.ValueOf().MethodByName` to get the reflect.method variable
4. Assemble params and call the function `reflect.Method.Call()` to return returns

### API functions definition

```go
type EthInterface interface {
// EthGetBlockByHash Returns information about a block by hash.
// Parameter Name
// BlockHash:32 byte hex value
// HydratedTransactions:hydrated
// Result
// Return: Block
EthGetBlockByHash(BlockHash string, HydratedTransactions bool) (result Block, err error)
// EthGetBlockByNumber Returns information about a block by number.
// Parameter Name
// BlockNumber:hex encoded unsigned integer
// HydratedTransactions:hydrated
// Result
// Return: Block
EthGetBlockByNumber(BlockNumber string, HydratedTransactions bool) (result Block, err error)
// EthGetBlockTransactionCountByHash Returns the number of transactions in a block from a block matching the given block hash.
// Parameter Name
// BlockHash:32 byte hex value
// Result
// Return: Transaction count[hex encoded unsigned integer]
EthGetBlockTransactionCountByHash(BlockHash string) (result []string, err error)
// EthGetBlockTransactionCountByNumber Returns the number of transactions in a block matching the given block number.
// Parameter Name
// BlockNumber:hex encoded unsigned integer
// Result
// Return: Transaction count[hex encoded unsigned integer]
EthGetBlockTransactionCountByNumber(BlockNumber string) (result []string, err error)
// EthGetUncleCountByBlockHash Returns the number of uncles in a block from a block matching the given block hash.
// Parameter Name
// BlockHash:32 byte hex value
// Result
// Return: Uncle count[hex encoded unsigned integer]
EthGetUncleCountByBlockHash(BlockHash string) (result []string, err error)
// EthGetUncleCountByBlockNumber Returns the number of transactions in a block matching the given block number.
// Parameter Name
// BlockNumber:hex encoded unsigned integer
// Result
// Return: Uncle count[hex encoded unsigned integer]
EthGetUncleCountByBlockNumber(BlockNumber string) (result []string, err error)
// EthProtocolVersion Returns the current ethereum protocol version.
// Parameter Name
// ---
// Result
// Return: version
EthProtocolVersion() (result string, err error)
// EthSyncing Returns an object with data about the sync status or false.
// Parameter Name
// ---
// Result
// Return: Syncing status
EthSyncing() (result Syncing, err error)
// EthCoinbase Returns the client coinbase address.
// Parameter Name
// ---
// Result
// Return: Coinbase address (hex encoded address)
EthCoinbase() (result string, err error)
// EthAccounts Returns a list of addresses owned by client.
// Parameter Name
// ---
// Returns: Accounts[hex encoded address]
EthAccounts() (result []string, err error)
// EthBlockNumber Returns the number of most recent block.
// Parameter Name
// ---
// Result
// Return: Block number (hex encoded unsigned integer)
EthBlockNumber() (result string, err error)
// EthCall Executes a new message call immediately without creating a transaction on the block chain.
// Parameter Name
// Transaction:Transaction object with sender
// Result
// Return: Return data (hex encoded bytes)
EthCall(Transaction TransactionObjectWithSender) (result string, err error)
// EthEstimateGas Generates and returns an estimate of how much gas is necessary to allow the transaction to complete.
// Parameter Name
// Transaction:Transaction object with sender
// Result
// Return: Gas used (hex encoded unsigned integer)
EthEstimateGas(Transaction TransactionObjectWithSender) (result string, err error)
// EthGasPrice Returns the current price per gas in wei.
// Parameter Name
// ---
// Result
// Return: Gas price
EthGasPrice() (result string, err error)
// EthFeeHistory
// Parameter Name
// BlockCount:hex encoded unsigned integer
// NewestBlock:Block number or tag
// RewardPercentiles:rewardPercentiles
// Result
// Return: feeHistoryResult (Fee history for the returned block range. This can be a subsection of the requested range if not all blocks are available.)
EthFeeHistory(BlockCount string, NewestBlock string, RewardPercentiles []int) (result EthFeeHistoryResult, err error)
// EthNewFilter Creates a filter object, based on filter options, to notify when the state changes (logs).
// Parameter Name
// Filter:filter
// Return: Filter Identifier (hex encoded unsigned integer)
EthNewFilter(Filter Filter) (result string, err error)
// EthNewBlockFilter Creates a filter in the node, to notify when a new block arrives.
// Parameter Name
// ---
// Result
// Return: Filter Identifier (hex encoded unsigned integer)
EthNewBlockFilter() (result string, err error)
// EthNewPendingTransactionFilter Creates a filter in the node, to notify when new pending transactions arrive.
// Parameter Name
// ---
// Result
// Return: Filter Identifier (hex encoded unsigned integer)
EthNewPendingTransactionFilter() (result string, err error)
// EthUninstallFilter Uninstalls a filter with given id.
// Parameter Name
// FilterIdentifier: Filter Identifier (hex encoded unsigned integer)
// Result
// Return: Success(boolean)
EthUninstallFilter(FilterIdentifier string) (result bool, err error)
// EthGetFilterChanges Polling method for a filter, which returns an array of logs which occurred since last poll.
// Parameter Name
// FilterIdentifier:hex encoded unsigned integer
// Result
// Return: Log objects [ LogResult ]
EthGetFilterChanges(FilterIdentifier string) (result LogResult, err error)
// EthGetFilterLogs Returns an array of all logs matching filter with given id.
// Parameter Name
// FilterIdentifier:hex encoded unsigned integer
// Result
// Return: Log objects [ LogResult ]
EthGetFilterLogs(FilterIdentifier string) (result LogResult, err error)
// EthGetLogs Returns an array of all logs matching filter with given id.
// Parameter Name
// Filter:filter
// Result
// Return: Log objects [ LogResult ]
EthGetLogs(Filter Filter) (result LogResult, err error)
// EthMining Returns whether the client is actively mining new blocks.
// Parameter Name
// Result
// Return: miningStatus(boolean)
EthMining() (result bool, err error)
// EthHashrate Returns the number of hashes per second that the node is mining with.
// Parameter Name
// ---
// Result
// Return: miningStatus (Hashrate)
EthHashrate() (result string, err error)
// EthGetWork Returns the hash of the current block, the seedHash, and the boundary condition to be met (“target”).
// Parameter Name
// ---
// Result
// Return: Current work [Proof-of-work hash]
EthGetWork() (result []string, err error)
// EthSubmitWork Used for submitting a proof-of-work solution.
// Parameter Name
// ProofOfWorkHash:32 hex encoded bytes
// SeedHash:32 hex encoded bytes
// Difficulty:32 hex encoded bytes
// Result
// Return: Success (boolean)
EthSubmitWork(ProofOfWorkHash string, SeedHash string, Difficulty string) (result bool, err error)
// EthSubmitHashrate Used for submitting mining hashrate.
// Parameter Name
// Hashrate:32 hex encoded bytes
// ID:32 hex encoded bytes
// Result
// Return: Success(boolean)
EthSubmitHashrate(Hashrate string, ID string) (result bool, err error)
// EthSign Returns an EIP-191 signature over the provided data.
// Parameter
// Address:hex encoded address
// Message:hex encoded bytes
// Result
// Return: Signature(65 hex encoded bytes)
EthSign(Address string, Message string) (result string, err error)
// EthSignTransaction Returns an RLP encoded transaction signed by the specified account.
// Parameter Name
// Transaction:Transaction object with sender
// Result
// Return: Encoded transaction (hex encoded bytes)
EthSignTransaction(Transaction TransactionObjectWithSender) (result string, err error)
// EthGetBalance Returns the balance of the account of given address.
// Parameter Name
// Address:hex encoded address
// Block:Block number or tag
// Result
// Return: Balance (hex encoded unsigned integer)
EthGetBalance(Address string, Block string) (result string, err error)
// EthGetStorage Returns the value from a storage position at a given address.
// Parameter Name
// Address:hex encoded address
// StorageSlot:hex encoded unsigned integer
// Block:Block number or tag
// Result
// Return: Value (hex encoded bytes)
EthGetStorage(Address string, StorageSlot string, Block string) (result string, err error)
// EthGetTransactionCount Returns the number of transactions sent from an address.
// Address:hex encoded address
// Block:Block number or tag
// Return: Transaction count[hex encoded unsigned integer]
EthGetTransactionCount(Address string, Block string) (result []string, err error)
// EthGetCode Returns code at a given address.
// Parameter Name
// Address:hex encoded address
// Block:Block number or tag
// Result
// Return: Bytecode(hex encoded bytes)
EthGetCode(Address string, Block string) (result string, err error)
// EthSendTransaction Signs and submits a transaction.
// Parameter Name
// Transaction:Transaction object with sender
// Result
// Return: 32 byte hex value
EthSendTransaction(Transaction TransactionObjectWithSender) (result string, err error)
// EthSendRawTransaction Submits a raw transaction.
// Parameter Name
// Transaction:hex encoded bytes
// Result
// Return: Transaction hash(32 byte hex value)
EthSendRawTransaction(Transaction string) (result string, err error)
// EthGetTransactionByHash Returns the information about a transaction requested by transaction hash.
// Parameter Name
// TransactionHash:32 byte hex value
// Result
// Return: Transaction information
EthGetTransactionByHash(TransactionHash string) (result TransactionInformation, err error)
// EthGetTransactionByBlockHashAndIndex Returns information about a transaction by block hash and transaction index position.
// Parameter Name
// BlockHash:32 byte hex value
// TransactionIndex:hex encoded unsigned integer
// Result
// Return: Transaction information
EthGetTransactionByBlockHashAndIndex(BlockHash string, TransactionIndex string) (result TransactionInformation, err error)
// EthGetTransactionByBlockNumberAndIndex Returns information about a transaction by block number and transaction index position.
// Parameter Name
// BlockNumber:hex encoded unsigned integer
// TransactionIndex:hex encoded unsigned integer
// Result
// Return: Transaction information
EthGetTransactionByBlockNumberAndIndex(BlockNumber string, TransactionIndex string) (result TransactionInformation, err error)
// EthGetTransactionReceipt Returns the receipt of a transaction by transaction hash.
// Parameter Name
// TransactionHash:32 byte hex value
// Result
// Return: Receipt Information (EthGetTransactionReceiptResult)
EthGetTransactionReceipt(TransactionHash string) (result EthGetTransactionReceiptResult, err error)
}
```