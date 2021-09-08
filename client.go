package SealRPC

import (
	"fmt"
	"github.com/ybbus/jsonrpc/v2"
)

type EthRPCClient struct {
	c jsonrpc.RPCClient
}

func NewEthRPCClient(rpcURL string) *EthRPCClient {
	rpcClient := jsonrpc.NewClient(rpcURL)
	return &EthRPCClient{c: rpcClient}
}

// EthGetBlockByHash Returns information about a block by hash.
// 按哈希返回有关块的信息。
// BlockHash:32 byte hex value
// HydratedTransactions:hydrated
func (e *EthRPCClient) EthGetBlockByHash(BlockHash string, HydratedTransactions bool) (result Block, err error) {
	err = e.c.CallFor(&result, "eth_getBlockByHash", BlockHash, HydratedTransactions)
	return
}

// EthGetBlockByNumber Returns information about a block by number.
// 按编号返回有关块的信息。
// BlockNumber:hex encoded unsigned integer
// HydratedTransactions:hydrated
func (e *EthRPCClient) EthGetBlockByNumber(BlockNumber string, HydratedTransactions bool) (result Block, err error) {
	err = e.c.CallFor(&result, "eth_getBlockByNumber", BlockNumber, HydratedTransactions)
	return
}

// EthGetBlockTransactionCountByHash Returns the number of transactions in a block from a block matching the given block hash.
// 从匹配给定块哈希块返回块交易的数量。
// BlockHash:32 byte hex value
// Return: Transaction count[hex encoded unsigned integer]
func (e *EthRPCClient) EthGetBlockTransactionCountByHash(BlockHash string) (result []string, err error) {
	err = e.c.CallFor(&result, "eth_getBlockTransactionCountByHash", BlockHash)
	return
}

// EthGetBlockTransactionCountByNumber Returns the number of transactions in a block matching the given block number.
// 返回块中与给定块号匹配的交易数。
// BlockNumber:hex encoded unsigned integer
// Return: Transaction count[hex encoded unsigned integer]
func (e *EthRPCClient) EthGetBlockTransactionCountByNumber(BlockNumber string) (result []string, err error) {
	err = e.c.CallFor(&result, "eth_getBlockTransactionCountByNumber", BlockNumber)
	return
}

// EthGetUncleCountByBlockHash Returns the number of uncles in a block from a block matching the given block hash.
// 从与给定块哈希匹配的块返回块中的uncles。
// BlockHash:32 byte hex value
// Return: Uncle count[hex encoded unsigned integer]
func (e *EthRPCClient) EthGetUncleCountByBlockHash(BlockHash string) (result []string, err error) {
	err = e.c.CallFor(&result, "eth_getUncleCountByBlockHash", BlockHash)
	return
}

// EthGetUncleCountByBlockNumber Returns the number of transactions in a block matching the given block number.
// 返回块中与给定块号匹配的交易数。
// BlockNumber:hex encoded unsigned integer
// Return: Uncle count[hex encoded unsigned integer]
func (e *EthRPCClient) EthGetUncleCountByBlockNumber(BlockNumber string) (result []string, err error) {
	err = e.c.CallFor(&result, "eth_getUncleCountByBlockNumber", BlockNumber)
	return
}

// EthProtocolVersion Returns the current ethereum protocol version.
// 返回当前以太坊协议版本。
// Return: version
func (e *EthRPCClient) EthProtocolVersion() (result string, err error) {
	err = e.c.CallFor(&result, "eth_protocolVersion")
	return
}

// EthSyncing Returns an object with data about the sync status or false.
// 返回一个对象，其中包含有关同步状态的数据或false。
func (e *EthRPCClient) EthSyncing() (result Syncing, err error) {
	err = e.c.CallFor(&result, "eth_syncing")
	return
}

// EthCoinbase Returns the client coinbase address.
// 返回客户端coinbase地址。
// Return: hex encoded address
func (e *EthRPCClient) EthCoinbase() (result string, err error) {
	err = e.c.CallFor(&result, "eth_coinbase")
	return
}

// EthAccounts Returns a list of addresses owned by client.
// 返回客户端拥有的地址列表。
// Returns: Accounts[hex encoded address]
func (e *EthRPCClient) EthAccounts() (result []string, err error) {
	err = e.c.CallFor(&result, "eth_accounts")
	return
}

// EthBlockNumber Returns the number of most recent block.
// 返回最近的块数。
// Return: hex encoded unsigned integer
func (e *EthRPCClient) EthBlockNumber() (result string, err error) {
	err = e.c.CallFor(&result, "eth_blockNumber")
	return
}

// EthCall Executes a new message call immediately without creating a transaction on the block chain.
// 立即执行新的消息调用，而不在块链上创建事务。
// Transaction:Transaction object with sender
// Return: hex encoded bytes
func (e *EthRPCClient) EthCall(Transaction TransactionObjectWithSender) (result string, err error) {
	err = e.c.CallFor(&result, "eth_call", Transaction)
	return
}

// EthEstimateGas Generates and returns an estimate of how much gas is necessary to allow the transaction to complete.
// 生成并返回允许交易完成所需的天然气量的估计值。
// Transaction:Transaction object with sender
// Return: hex encoded unsigned integer
func (e *EthRPCClient) EthEstimateGas(Transaction TransactionObjectWithSender) (result string, err error) {
	err = e.c.CallFor(&result, "eth_estimateGas", Transaction)
	return
}

// EthGasPrice Returns the current price per gas in wei.
// 返回每种气体的当前价格（单位：wei）。
// Return: Gas price
func (e *EthRPCClient) EthGasPrice() (result string, err error) {
	err = e.c.CallFor(&result, "eth_gasPrice")
	return
}

// EthFeeHistory
//
// BlockCount:hex encoded unsigned integer
// NewestBlock:Block number or tag
// RewardPercentiles:rewardPercentiles
func (e *EthRPCClient) EthFeeHistory(BlockCount string, NewestBlock string, RewardPercentiles []int) (result EthFeeHistoryResult, err error) {
	err = e.c.CallFor(&result, "eth_feeHistory", BlockCount, NewestBlock, RewardPercentiles)
	return
}

// EthNewFilter Creates a filter object, based on filter options, to notify when the state changes (logs).
// 根据过滤器选项创建过滤器对象，以在状态更改（日志）时发出通知。
// Filter:filter
// Return: hex encoded unsigned integer
func (e *EthRPCClient) EthNewFilter(Filter Filter) (result string, err error) {
	err = e.c.CallFor(&result, "eth_newFilter", Filter)
	return
}

// EthNewBlockFilter Creates a filter in the node, to notify when a new block arrives.
// 在节点中创建过滤器，以在新块到达时通知。
// Return: hex encoded unsigned integer
func (e *EthRPCClient) EthNewBlockFilter() (result string, err error) {
	err = e.c.CallFor(&result, "eth_newBlockFilter")
	return
}

// EthNewPendingTransactionFilter Creates a filter in the node, to notify when new pending transactions arrive.
// 在节点中创建筛选器，以便在新的挂起事务到达时通知。
// Return: hex encoded unsigned integer
func (e *EthRPCClient) EthNewPendingTransactionFilter() (result string, err error) {
	err = e.c.CallFor(&result, "eth_newPendingTransactionFilter")
	return
}

// EthUninstallFilter Uninstalls a filter with given id.
// 卸载具有给定id的筛选器。
// FilterIdentifier:hex encoded unsigned integer
func (e *EthRPCClient) EthUninstallFilter(FilterIdentifier string) (result bool, err error) {
	err = e.c.CallFor(&result, "eth_uninstallFilter", FilterIdentifier)
	return
}

// EthGetFilterChanges Polling method for a filter, which returns an array of logs which occurred since last poll.
// 筛选器的轮询方法，该方法返回自上次轮询以来发生的日志数组。
// FilterIdentifier:hex encoded unsigned integer
func (e *EthRPCClient) EthGetFilterChanges(FilterIdentifier string) (result LogResult, err error) {
	err = e.c.CallFor(&result, "eth_getFilterChanges", FilterIdentifier)
	return
}

// EthGetFilterLogs Returns an array of all logs matching filter with given id.
// 返回与给定id的筛选器匹配的所有日志的数组。
// FilterIdentifier:hex encoded unsigned integer
func (e *EthRPCClient) EthGetFilterLogs(FilterIdentifier string) (result LogResult, err error) {
	err = e.c.CallFor(&result, "eth_getFilterLogs", FilterIdentifier)
	return
}

// EthGetLogs Returns an array of all logs matching filter with given id.
// 返回与给定id的筛选器匹配的所有日志的数组。
// Filter:filter
func (e *EthRPCClient) EthGetLogs(Filter Filter) (result LogResult, err error) {
	err = e.c.CallFor(&result, "eth_getLogs", Filter)
	return
}

// EthMining Returns whether the client is actively mining new blocks.
// 返回客户端是否正在积极挖掘新块
// Return: miningStatus
func (e *EthRPCClient) EthMining() (result bool, err error) {
	err = e.c.CallFor(&result, "eth_mining")
	return
}

// EthHashrate Returns the number of hashes per second that the node is mining with.
// 返回节点每秒使用的哈希数。
// Return: Hashrate
func (e *EthRPCClient) EthHashrate() (result string, err error) {
	err = e.c.CallFor(&result, "eth_hashrate")
	return
}

// EthGetWork Returns the hash of the current block, the seedHash, and the boundary condition to be met (“target”).
// 返回当前块的散列、种子散列和要满足的边界条件（“目标”）。
// Return: Proof-of-work hash
func (e *EthRPCClient) EthGetWork() (result []string, err error) {
	err = e.c.CallFor(&result, "eth_getWork")
	return
}

// EthSubmitWork Used for submitting a proof-of-work solution.
// 用于提交工作证明解决方案。
// ProofOfWorkHash:32 hex encoded bytes
// SeedHash:32 hex encoded bytes
// Difficulty:32 hex encoded bytes
func (e *EthRPCClient) EthSubmitWork(ProofOfWorkHash string, SeedHash string, Difficulty string) (result bool, err error) {
	err = e.c.CallFor(&result, "eth_submitWork", ProofOfWorkHash, SeedHash, Difficulty)
	return
}

// EthSubmitHashrate Used for submitting mining hashrate.
// 用于提交挖掘哈希率。
// Hashrate:32 hex encoded bytes
// ID:32 hex encoded bytes
func (e *EthRPCClient) EthSubmitHashrate(Hashrate string, ID string) (result bool, err error) {
	err = e.c.CallFor(&result, "eth_submitHashrate", Hashrate, ID)
	return
}

// EthSign Returns an EIP-191 signature over the provided data.
// 返回提供数据上的EIP-191签名。
// Address:hex encoded address
// Message:hex encoded bytes
// Return: 65 hex encoded bytes
func (e *EthRPCClient) EthSign(Address string, Message string) (result string, err error) {
	err = e.c.CallFor(&result, "eth_sign", Address, Message)
	return
}

// EthSignTransaction Returns an RLP encoded transaction signed by the specified account.
// 返回由指定帐户签名的RLP编码事务。
// Transaction:Transaction object with sender
// Return: hex encoded bytes
func (e *EthRPCClient) EthSignTransaction(Transaction TransactionObjectWithSender) (result string, err error) {
	err = e.c.CallFor(&result, "eth_signTransaction", Transaction)
	return
}

// EthGetBalance Returns the balance of the account of given address.
// 返回给定地址的帐户余额
// Address:hex encoded address
// Block:Block number or tag
// Return: hex encoded unsigned integer
func (e *EthRPCClient) EthGetBalance(Address string, Block string) (result string, err error) {
	err = e.c.CallFor(&result, "eth_getBalance", Address, Block)
	return
}

// EthGetStorage Returns the value from a storage position at a given address.
// 从给定地址的存储位置返回值。
// Address:hex encoded address
// StorageSlot:hex encoded unsigned integer
// Block:Block number or tag
// Return: hex encoded bytes
func (e *EthRPCClient) EthGetStorage(Address string, StorageSlot string, Block string) (result string, err error) {
	err = e.c.CallFor(&result, "eth_getStorage", Address, StorageSlot, Block)
	return
}

// EthGetTransactionCount Returns the number of transactions sent from an address.
// 返回从某个地址发送的事务数。
// Address:hex encoded address
// Block:Block number or tag
// Return: Transaction count[hex encoded unsigned integer]
func (e *EthRPCClient) EthGetTransactionCount(Address string, Block string) (result []string, err error) {
	err = e.c.CallFor(&result, "eth_getTransactionCount", Address, Block)
	return
}

// EthGetCode Returns code at a given address.
// 返回给定地址处的代码。
// Address:hex encoded address
// Block:Block number or tag
// Return: hex encoded bytes
func (e *EthRPCClient) EthGetCode(Address string, Block string) (result string, err error) {
	err = e.c.CallFor(&result, "eth_getCode", Address, Block)
	return
}

// EthSendTransaction Signs and submits a transaction.
// 签署并提交交易。
// Transaction:Transaction object with sender
// Return: 32 byte hex value
func (e *EthRPCClient) EthSendTransaction(Transaction TransactionObjectWithSender) (result string, err error) {
	err = e.c.CallFor(&result, "eth_sendTransaction", Transaction)
	return
}

// EthSendRawTransaction Submits a raw transaction.
// 提交未处理的交易。
// Transaction:hex encoded bytes
// Return: 32 byte hex value
func (e *EthRPCClient) EthSendRawTransaction(Transaction string) (result string, err error) {
	err = e.c.CallFor(&result, "eth_sendRawTransaction", Transaction)
	return
}

// EthGetTransactionByHash Returns the information about a transaction requested by transaction hash.
// 返回有关事务哈希请求的事务的信息。
// TransactionHash:32 byte hex value
// Return: Transaction information
func (e *EthRPCClient) EthGetTransactionByHash(TransactionHash string) (result TransactionInformation, err error) {
	err = e.c.CallFor(&result, "eth_getTransactionByHash", TransactionHash)
	return
}

// EthGetTransactionByBlockHashAndIndex Returns information about a transaction by block hash and transaction index position.
// 按块哈希和事务索引位置返回有关交易的信息。
// BlockHash:32 byte hex value
// TransactionIndex:hex encoded unsigned integer
// Return: Transaction information
func (e *EthRPCClient) EthGetTransactionByBlockHashAndIndex(BlockHash string, TransactionIndex string) (result TransactionInformation, err error) {
	err = e.c.CallFor(&result, "eth_getTransactionByBlockHashAndIndex", BlockHash, TransactionIndex)
	return
}

// EthGetTransactionByBlockNumberAndIndex Returns information about a transaction by block number and transaction index position.
// 按块号和事务索引位置返回有关事务的信息。
// BlockNumber:hex encoded unsigned integer
// TransactionIndex:hex encoded unsigned integer
// Return: Transaction information
func (e *EthRPCClient) EthGetTransactionByBlockNumberAndIndex(BlockNumber string, TransactionIndex string) (result TransactionInformation, err error) {
	err = e.c.CallFor(&result, "eth_getTransactionByBlockNumberAndIndex", BlockNumber, TransactionIndex)
	return
}

// EthGetTransactionReceipt Returns the receipt of a transaction by transaction hash.
// 按事务哈希返回事务的收据。
// TransactionHash:32 byte hex value
func (e *EthRPCClient) EthGetTransactionReceipt(TransactionHash string) (result EthGetTransactionReceiptResult, err error) {
	err = e.c.CallFor(&result, "eth_getTransactionReceipt", TransactionHash)
	return
}

func main() {
	client := NewEthRPCClient("http://127.0.0.1:8545")
	number, err := client.EthGetBlockByNumber("0x1", true)
	fmt.Println(number, err)
	index, err := client.EthGetBlockTransactionCountByHash("0xe9dc52ef255d3b8ee11b35acc4b25f41ce096c71392aafe58907df9e1dbaaa4b")
	fmt.Println(index, err)
}

type Server struct {
	methodName string
	params     []interface{}
	resultType interface{}
}

type evm struct {
	//返回当前以太坊协议版本。
	//返回一个对象，其中包含有关同步状态的数据或false。
	//返回客户端coinbase地址。
	//返回客户端拥有的地址列表。
	//返回最近的块数。
	//立即执行新的消息调用，而不在块链上创建事务。(执行交易)
	gaps struct {
		//生成并返回允许交易完成所需的天然气量的估计值。
		//返回每种气体的当前价格（单位：wei）。
		//小费历史数据
	}
	//根据过滤器选项创建过滤器对象，以在状态更改（日志）时发出通知。
	//在节点中创建过滤器，以在新块到达时通知。
	//在节点中创建筛选器，以便在新的挂起事务到达时通知。
	//卸载具有给定id的筛选器。
	filter struct {
		//筛选器的轮询方法，该方法返回自上次轮询以来发生的日志数组。
		//返回与给定id的筛选器匹配的所有日志的数组。
		//返回与给定id的筛选器匹配的所有日志的数组。
	}
	//返回客户端是否正在积极挖掘新块
	//返回节点每秒使用的哈希数。
	//返回当前块的散列、种子散列和要满足的边界条件（“目标”）。
	//用于提交工作证明解决方案。
	//用于提交挖掘哈希率。

	Address struct {
		//返回提供数据上的EIP-191签名。
		//返回给定地址的帐户余额
		//从给定地址的存储位置返回值。
		//返回从某个地址发送的事务数。
		//返回给定地址处的代码。
	}
	Block struct {
		//按哈希返回有关块的信息。
		//按块的编号返回有关块的信息。

		//签署并提交交易。
		//提交未处理的交易。

		//返回块中与给定块号匹配的交易数。
		//返回块中与给定块号匹配的交易数。
		//从匹配给定块哈希块返回块交易的数量。
		Transaction struct {
			//返回由指定帐户签名的RLP编码事务。
			//返回有关事务哈希请求的事务的信息。
			//按块哈希和事务索引位置返回有关交易的信息。
			//按块号和事务索引位置返回有关事务的信息。
			//按事务哈希返回事务的收据。
		}

		//从与给定块哈希匹配的块返回块中的uncles。
		uncles struct {
		}
	}
}
