package SealRPC

import (
	"github.com/ybbus/jsonrpc/v2"
)

type JsonRPCClient struct {
	c jsonrpc.RPCClient
}

//NewJsonRPCClient Create a NewJsonRPCClient for JsonRPCURL
func NewJsonRPCClient(rpcURL string) *JsonRPCClient {
	rpcClient := jsonrpc.NewClient(rpcURL)
	return &JsonRPCClient{c: rpcClient}
}

func (e *JsonRPCClient) EthGetBlockByHash(BlockHash string, HydratedTransactions bool) (result Block, err error) {
	err = e.c.CallFor(&result, "eth_getBlockByHash", BlockHash, HydratedTransactions)
	return
}

func (e *JsonRPCClient) EthGetBlockByNumber(BlockNumber string, HydratedTransactions bool) (result Block, err error) {
	err = e.c.CallFor(&result, "eth_getBlockByNumber", BlockNumber, HydratedTransactions)
	return
}

func (e *JsonRPCClient) EthGetBlockTransactionCountByHash(BlockHash string) (result []string, err error) {
	err = e.c.CallFor(&result, "eth_getBlockTransactionCountByHash", BlockHash)
	return
}

func (e *JsonRPCClient) EthGetBlockTransactionCountByNumber(BlockNumber string) (result []string, err error) {
	err = e.c.CallFor(&result, "eth_getBlockTransactionCountByNumber", BlockNumber)
	return
}

func (e *JsonRPCClient) EthGetUncleCountByBlockHash(BlockHash string) (result []string, err error) {
	err = e.c.CallFor(&result, "eth_getUncleCountByBlockHash", BlockHash)
	return
}

func (e *JsonRPCClient) EthGetUncleCountByBlockNumber(BlockNumber string) (result []string, err error) {
	err = e.c.CallFor(&result, "eth_getUncleCountByBlockNumber", BlockNumber)
	return
}

func (e *JsonRPCClient) EthProtocolVersion() (result string, err error) {
	err = e.c.CallFor(&result, "eth_protocolVersion")
	return
}

func (e *JsonRPCClient) EthSyncing() (result Syncing, err error) {
	err = e.c.CallFor(&result, "eth_syncing")
	return
}

func (e *JsonRPCClient) EthCoinbase() (result string, err error) {
	err = e.c.CallFor(&result, "eth_coinbase")
	return
}

func (e *JsonRPCClient) EthAccounts() (result []string, err error) {
	err = e.c.CallFor(&result, "eth_accounts")
	return
}

func (e *JsonRPCClient) EthBlockNumber() (result string, err error) {
	err = e.c.CallFor(&result, "eth_blockNumber")
	return
}

func (e *JsonRPCClient) EthCall(Transaction TransactionObjectWithSender) (result string, err error) {
	err = e.c.CallFor(&result, "eth_call", Transaction)
	return
}

func (e *JsonRPCClient) EthEstimateGas(Transaction TransactionObjectWithSender) (result string, err error) {
	err = e.c.CallFor(&result, "eth_estimateGas", Transaction)
	return
}

func (e *JsonRPCClient) EthGasPrice() (result string, err error) {
	err = e.c.CallFor(&result, "eth_gasPrice")
	return
}

func (e *JsonRPCClient) EthFeeHistory(BlockCount string, NewestBlock string, RewardPercentiles []int) (result EthFeeHistoryResult, err error) {
	err = e.c.CallFor(&result, "eth_feeHistory", BlockCount, NewestBlock, RewardPercentiles)
	return
}

func (e *JsonRPCClient) EthNewFilter(Filter Filter) (result string, err error) {
	err = e.c.CallFor(&result, "eth_newFilter", Filter)
	return
}

func (e *JsonRPCClient) EthNewBlockFilter() (result string, err error) {
	err = e.c.CallFor(&result, "eth_newBlockFilter")
	return
}

func (e *JsonRPCClient) EthNewPendingTransactionFilter() (result string, err error) {
	err = e.c.CallFor(&result, "eth_newPendingTransactionFilter")
	return
}

func (e *JsonRPCClient) EthUninstallFilter(FilterIdentifier string) (result bool, err error) {
	err = e.c.CallFor(&result, "eth_uninstallFilter", FilterIdentifier)
	return
}

func (e *JsonRPCClient) EthGetFilterChanges(FilterIdentifier string) (result LogResult, err error) {
	err = e.c.CallFor(&result, "eth_getFilterChanges", FilterIdentifier)
	return
}

func (e *JsonRPCClient) EthGetFilterLogs(FilterIdentifier string) (result LogResult, err error) {
	err = e.c.CallFor(&result, "eth_getFilterLogs", FilterIdentifier)
	return
}

func (e *JsonRPCClient) EthGetLogs(Filter Filter) (result LogResult, err error) {
	err = e.c.CallFor(&result, "eth_getLogs", Filter)
	return
}

func (e *JsonRPCClient) EthMining() (result bool, err error) {
	err = e.c.CallFor(&result, "eth_mining")
	return
}

func (e *JsonRPCClient) EthHashrate() (result string, err error) {
	err = e.c.CallFor(&result, "eth_hashrate")
	return
}

func (e *JsonRPCClient) EthGetWork() (result []string, err error) {
	err = e.c.CallFor(&result, "eth_getWork")
	return
}

func (e *JsonRPCClient) EthSubmitWork(ProofOfWorkHash string, SeedHash string, Difficulty string) (result bool, err error) {
	err = e.c.CallFor(&result, "eth_submitWork", ProofOfWorkHash, SeedHash, Difficulty)
	return
}

func (e *JsonRPCClient) EthSubmitHashrate(Hashrate string, ID string) (result bool, err error) {
	err = e.c.CallFor(&result, "eth_submitHashrate", Hashrate, ID)
	return
}

func (e *JsonRPCClient) EthSign(Address string, Message string) (result string, err error) {
	err = e.c.CallFor(&result, "eth_sign", Address, Message)
	return
}

func (e *JsonRPCClient) EthSignTransaction(Transaction TransactionObjectWithSender) (result string, err error) {
	err = e.c.CallFor(&result, "eth_signTransaction", Transaction)
	return
}

func (e *JsonRPCClient) EthGetBalance(Address string, Block string) (result string, err error) {
	err = e.c.CallFor(&result, "eth_getBalance", Address, Block)
	return
}

func (e *JsonRPCClient) EthGetStorage(Address string, StorageSlot string, Block string) (result string, err error) {
	err = e.c.CallFor(&result, "eth_getStorage", Address, StorageSlot, Block)
	return
}

func (e *JsonRPCClient) EthGetTransactionCount(Address string, Block string) (result []string, err error) {
	err = e.c.CallFor(&result, "eth_getTransactionCount", Address, Block)
	return
}

func (e *JsonRPCClient) EthGetCode(Address string, Block string) (result string, err error) {
	err = e.c.CallFor(&result, "eth_getCode", Address, Block)
	return
}

func (e *JsonRPCClient) EthSendTransaction(Transaction TransactionObjectWithSender) (result string, err error) {
	err = e.c.CallFor(&result, "eth_sendTransaction", Transaction)
	return
}

func (e *JsonRPCClient) EthSendRawTransaction(Transaction string) (result string, err error) {
	err = e.c.CallFor(&result, "eth_sendRawTransaction", Transaction)
	return
}

func (e *JsonRPCClient) EthGetTransactionByHash(TransactionHash string) (result TransactionInformation, err error) {
	err = e.c.CallFor(&result, "eth_getTransactionByHash", TransactionHash)
	return
}

func (e *JsonRPCClient) EthGetTransactionByBlockHashAndIndex(BlockHash string, TransactionIndex string) (result TransactionInformation, err error) {
	err = e.c.CallFor(&result, "eth_getTransactionByBlockHashAndIndex", BlockHash, TransactionIndex)
	return
}

func (e *JsonRPCClient) EthGetTransactionByBlockNumberAndIndex(BlockNumber string, TransactionIndex string) (result TransactionInformation, err error) {
	err = e.c.CallFor(&result, "eth_getTransactionByBlockNumberAndIndex", BlockNumber, TransactionIndex)
	return
}

func (e *JsonRPCClient) EthGetTransactionReceipt(TransactionHash string) (result EthGetTransactionReceiptResult, err error) {
	err = e.c.CallFor(&result, "eth_getTransactionReceipt", TransactionHash)
	return
}
