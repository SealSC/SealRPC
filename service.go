package SealRPC

import (
	"errors"
)

var UnregisteredErr = errors.New("function Unregistered")

type EthService struct {
	// EthGetBlockByHashFunc Returns information about a block by hash.
	// 按哈希返回有关块的信息。
	// BlockHash:32 byte hex value
	// HydratedTransactions:hydrated
	EthGetBlockByHashFuc func(BlockHash string, HydratedTransactions bool) (Block, error)
	// EthGetBlockByNumberFunc Returns information about a block by number.
	// 按编号返回有关块的信息。
	// BlockNumber:hex encoded unsigned integer
	// HydratedTransactions:hydrated
	EthGetBlockByNumberFuc func(BlockNumber string, HydratedTransactions bool) (Block, error)
	// EthGetBlockTransactionCountByHashFuc Returns the number of transactions in a block from a block matching the given block hash.
	// BlockHash:32 byte hex value
	EthGetBlockTransactionCountByHashFuc func(BlockHash string) ([]string, error)
	// EthGetBlockTransactionCountByNumberFunc Returns the number of transactions in a block matching the given block number.
	// 返回块中与给定块号匹配的交易数。
	// BlockNumber:hex encoded unsigned integer
	EthGetBlockTransactionCountByNumberFunc func(BlockNumber string) ([]string, error)
	// EthGetUncleCountByBlockHashFuncType Returns the number of uncles in a block from a block matching the given block hash.
	// 从与给定块哈希匹配的块返回块中的uncles。
	// BlockHash:32 byte hex value
	EthGetUncleCountByBlockHashFunc func(BlockHash string) ([]string, error)
	// EthGetUncleCountByBlockNumberFunc Returns the number of transactions in a block matching the given block number.
	// 返回块中与给定块号匹配的交易数。
	// BlockNumber:hex encoded unsigned integer
	EthGetUncleCountByBlockNumberFunc func(BlockNumber string) ([]string, error)
	// EthProtocolVersionFunc Returns the current ethereum protocol version.
	// 返回当前以太坊协议版本。
	EthProtocolVersionFunc func() (string, error)
	// EthSyncingFunc Returns an object with data about the sync status or false.
	// 返回一个对象，其中包含有关同步状态的数据或false。
	EthSyncingFunc func() (Syncing, error)
	// EthCoinbaseFunc Returns the client coinbase address.
	// 返回客户端coinbase地址。
	EthCoinbaseFunc func() (string, error)
	// EthAccountsFunc Returns a list of addresses owned by client.
	// 返回客户端拥有的地址列表。
	EthAccountsFunc func() ([]string, error)
	// EthBlockNumberFunc Returns the number of most recent block.
	// 返回最近的块数。
	EthBlockNumberFunc func() (string, error)
	// EthCallFunc Executes a new message call immediately without creating a transaction on the block chain.
	// 立即执行新的消息调用，而不在块链上创建事务。
	// Transaction:Transaction object with sender
	EthCallFunc func(Transaction TransactionObjectWithSender) (string, error)
	// EthEstimateGasFunc Generates and returns an estimate of how much gas is necessary to allow the transaction to complete.
	// 生成并返回允许交易完成所需的天然气量的估计值。
	// Transaction:Transaction object with sender
	EthEstimateGasFunc func(Transaction TransactionObjectWithSender) (string, error)
	// EthGasPriceFunc Returns the current price per gas in wei.
	// 返回每种气体的当前价格（单位：wei）。
	EthGasPriceFunc func() (string, error)
	// EthFeeHistoryFunc
	// BlockCount:hex encoded unsigned integer
	// NewestBlock:Block number or tag
	// RewardPercentiles:rewardPercentiles
	EthFeeHistoryFunc func(BlockCount string, NewestBlock string, RewardPercentiles []int) (EthFeeHistoryResult, error)
	// EthNewFilterFunc Creates a filter object, based on filter options, to notify when the state changes (logs).
	// 根据过滤器选项创建过滤器对象，以在状态更改（日志）时发出通知。
	// Filter:filter
	EthNewFilterFunc func(Filter Filter) (string, error)
	// EthNewBlockFilterFunc Creates a filter in the node, to notify when a new block arrives.
	// 在节点中创建过滤器，以在新块到达时通知。
	EthNewBlockFilterFunc func() (string, error)
	// EthNewPendingTransactionFilterFunc Creates a filter in the node, to notify when new pending transactions arrive.
	// 在节点中创建筛选器，以便在新的挂起事务到达时通知。
	EthNewPendingTransactionFilterFunc func() (string, error)
	// EthUninstallFilterFunc Uninstalls a filter with given id.
	// 卸载具有给定id的筛选器。
	// FilterIdentifier:hex encoded unsigned integer
	EthUninstallFilterFunc func(FilterIdentifier string) (bool, error)
	// EthGetFilterChangesFunc Polling method for a filter, which returns an array of logs which occurred since last poll.
	// 筛选器的轮询方法，该方法返回自上次轮询以来发生的日志数组。
	// FilterIdentifier:hex encoded unsigned integer
	EthGetFilterChangesFunc func(FilterIdentifier string) (LogResult, error)
	// EthGetFilterLogsFunc Returns an array of all logs matching filter with given id.
	// 返回与给定id的筛选器匹配的所有日志的数组。
	// FilterIdentifier:hex encoded unsigned integer
	EthGetFilterLogsFunc func(FilterIdentifier string) (LogResult, error)
	// EthGetLogsFunc Returns an array of all logs matching filter with given id.
	// 返回与给定id的筛选器匹配的所有日志的数组。
	// Filter:filter
	EthGetLogsFunc func(Filter Filter) (LogResult, error)
	// EthMiningFunc Returns whether the client is actively mining new blocks.
	// 返回客户端是否正在积极挖掘新块
	EthMiningFunc func() (bool, error)
	// EthHashrateFunc Returns the number of hashes per second that the node is mining with.
	// 返回节点每秒使用的哈希数。
	EthHashrateFunc func() (string, error)
	// EthGetWorkFunc Returns the hash of the current block, the seedHash, and the boundary condition to be met (“target”).
	// 返回当前块的散列、种子散列和要满足的边界条件（“目标”）。
	EthGetWorkFunc func() ([]string, error)
	// EthSubmitWorkFunc Used for submitting a proof-of-work solution.
	// 用于提交工作证明解决方案。
	// ProofOfWorkHash:32 hex encoded bytes
	// SeedHash:32 hex encoded bytes
	// Difficulty:32 hex encoded bytes
	EthSubmitWorkFunc func(ProofOfWorkHash string, SeedHash string, Difficulty string) (bool, error)
	// EthSubmitHashrateFunc Used for submitting mining hashrate.
	// 用于提交挖掘哈希率。
	// Hashrate:32 hex encoded bytes
	// ID:32 hex encoded bytes
	EthSubmitHashrateFunc func(Hashrate string, ID string) (bool, error)
	// EthSignFunc Returns an EIP-191 signature over the provided data.
	// 返回提供数据上的EIP-191签名。
	// Address:hex encoded address
	// Message:hex encoded bytes
	EthSignFunc func(Address string, Message string) (string, error)
	// EthSignTransactionFunc Returns an RLP encoded transaction signed by the specified account.
	// 返回由指定帐户签名的RLP编码事务。
	// Transaction:Transaction object with sender
	EthSignTransactionFunc func(Transaction TransactionObjectWithSender) (string, error)
	// EthGetBalanceFunc Returns the balance of the account of given address.
	// 返回给定地址的帐户余额
	// Address:hex encoded address
	// Block:Block number or tag
	EthGetBalanceFunc func(Address string, Block string) (string, error)
	// EthGetStorageFunc Returns the value from a storage position at a given address.
	// 从给定地址的存储位置返回值。
	// Address:hex encoded address
	// StorageSlot:hex encoded unsigned integer
	// Block:Block number or tag
	EthGetStorageFunc func(Address string, StorageSlot string, Block string) (string, error)
	// EthGetTransactionCountFunc Returns the number of transactions sent from an address.
	// 返回从某个地址发送的事务数。
	// Address:hex encoded address
	// Block:Block number or tag
	EthGetTransactionCountFunc func(Address string, Block string) ([]string, error)
	// EthGetCodeFunc Returns code at a given address.
	// 返回给定地址处的代码。
	// Address:hex encoded address
	// Block:Block number or tag
	EthGetCodeFunc func(Address string, Block string) (string, error)
	// EthSendTransactionFunc Signs and submits a transaction.
	// 签署并提交事务。
	// Transaction:Transaction object with sender
	EthSendTransactionFunc func(Transaction TransactionObjectWithSender) (string, error)
	// EthSendRawTransactionFunc Submits a raw transaction.
	// 提交未处理的事务。
	// Transaction:hex encoded bytes
	EthSendRawTransactionFunc func(Transaction string) (string, error)
	// EthGetTransactionByHashFunc Returns the information about a transaction requested by transaction hash.
	// 返回有关事务哈希请求的事务的信息。
	// TransactionHash:32 byte hex value
	EthGetTransactionByHashFunc func(TransactionHash string) (TransactionInformation, error)
	// EthGetTransactionByBlockHashAndIndexFunc Returns information about a transaction by block hash and transaction index position.
	// 按块哈希和事务索引位置返回有关事务的信息。
	// BlockHash:32 byte hex value
	// TransactionIndex:hex encoded unsigned integer
	EthGetTransactionByBlockHashAndIndexFunc func(BlockHash string, TransactionIndex string) (TransactionInformation, error)
	// EthGetTransactionByBlockNumberAndIndexFunc Returns information about a transaction by block number and transaction index position.
	// 按块号和事务索引位置返回有关事务的信息。
	// BlockNumber:hex encoded unsigned integer
	// TransactionIndex:hex encoded unsigned integer
	EthGetTransactionByBlockNumberAndIndexFunc func(BlockNumber string, TransactionIndex string) (TransactionInformation, error)
	// EthGetTransactionReceiptFunc Returns the receipt of a transaction by transaction hash.
	// 按事务哈希返回事务的收据。
	// TransactionHash:32 byte hex value
	EthGetTransactionReceiptFunc func(TransactionHash string) (EthGetTransactionReceiptResult, error)
}

func (e *EthService) EthGetBlockByHash(BlockHash string, HydratedTransactions bool) (result Block, err error) {
	if e.EthGetBlockByHashFuc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetBlockByHashFuc(BlockHash, HydratedTransactions)
}

func (e *EthService) EthGetBlockByNumber(BlockNumber string, HydratedTransactions bool) (result Block, err error) {
	if e.EthGetBlockByNumberFuc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetBlockByNumberFuc(BlockNumber, HydratedTransactions)
}

func (e *EthService) EthGetBlockTransactionCountByHash(BlockHash string) (result []string, err error) {
	if e.EthGetBlockTransactionCountByHashFuc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetBlockTransactionCountByHashFuc(BlockHash)
}

func (e *EthService) EthGetBlockTransactionCountByNumber(BlockNumber string) (result []string, err error) {
	if e.EthGetBlockTransactionCountByNumberFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetBlockTransactionCountByNumberFunc(BlockNumber)
}

func (e *EthService) EthGetUncleCountByBlockHash(BlockHash string) (result []string, err error) {
	if e.EthGetUncleCountByBlockHashFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetUncleCountByBlockHashFunc(BlockHash)
}

func (e *EthService) EthGetUncleCountByBlockNumber(BlockNumber string) (result []string, err error) {
	if e.EthGetUncleCountByBlockNumberFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetUncleCountByBlockNumberFunc(BlockNumber)
}

func (e *EthService) EthProtocolVersion() (result string, err error) {
	if e.EthProtocolVersionFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthProtocolVersionFunc()
}

func (e *EthService) EthSyncing() (result Syncing, err error) {
	if e.EthSyncingFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthSyncingFunc()
}

func (e *EthService) EthCoinbase() (result string, err error) {
	if e.EthCoinbaseFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthCoinbaseFunc()
}

func (e *EthService) EthAccounts() (result []string, err error) {
	if e.EthAccountsFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthAccountsFunc()
}

func (e *EthService) EthBlockNumber() (result string, err error) {
	if e.EthBlockNumberFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthBlockNumberFunc()
}

func (e *EthService) EthCall(Transaction TransactionObjectWithSender) (result string, err error) {
	if e.EthCallFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthCallFunc(Transaction)
}

func (e *EthService) EthEstimateGas(Transaction TransactionObjectWithSender) (result string, err error) {
	if e.EthEstimateGasFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthEstimateGasFunc(Transaction)
}

func (e *EthService) EthGasPrice() (result string, err error) {
	if e.EthGasPriceFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGasPriceFunc()
}

func (e *EthService) EthFeeHistory(BlockCount string, NewestBlock string, RewardPercentiles []int) (result EthFeeHistoryResult, err error) {
	if e.EthFeeHistoryFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthFeeHistoryFunc(BlockCount, NewestBlock, RewardPercentiles)
}

func (e *EthService) EthNewFilter(Filter Filter) (result string, err error) {
	if e.EthNewFilterFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthNewFilterFunc(Filter)
}

func (e *EthService) EthNewBlockFilter() (result string, err error) {
	if e.EthNewBlockFilterFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthNewBlockFilterFunc()
}

func (e *EthService) EthNewPendingTransactionFilter() (result string, err error) {
	if e.EthNewPendingTransactionFilterFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthNewPendingTransactionFilterFunc()
}

func (e *EthService) EthUninstallFilter(FilterIdentifier string) (result bool, err error) {
	if e.EthUninstallFilterFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthUninstallFilterFunc(FilterIdentifier)
}

func (e *EthService) EthGetFilterChanges(FilterIdentifier string) (result LogResult, err error) {
	if e.EthGetFilterChangesFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetFilterChangesFunc(FilterIdentifier)
}

func (e *EthService) EthGetFilterLogs(FilterIdentifier string) (result LogResult, err error) {
	if e.EthGetFilterLogsFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetFilterLogsFunc(FilterIdentifier)
}

func (e *EthService) EthGetLogs(Filter Filter) (result LogResult, err error) {
	if e.EthGetLogsFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetLogsFunc(Filter)
}

func (e *EthService) EthMining() (result bool, err error) {
	if e.EthMiningFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthMiningFunc()
}

func (e *EthService) EthHashrate() (result string, err error) {
	if e.EthHashrateFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthHashrateFunc()
}

func (e *EthService) EthGetWork() (result []string, err error) {
	if e.EthGetWorkFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetWorkFunc()
}

func (e *EthService) EthSubmitWork(ProofOfWorkHash string, SeedHash string, Difficulty string) (result bool, err error) {
	if e.EthSubmitWorkFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthSubmitWorkFunc(ProofOfWorkHash, SeedHash, Difficulty)
}

func (e *EthService) EthSubmitHashrate(Hashrate string, ID string) (result bool, err error) {
	if e.EthSubmitHashrateFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthSubmitHashrateFunc(Hashrate, ID)
}

func (e *EthService) EthSign(Address string, Message string) (result string, err error) {
	if e.EthSignFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthSignFunc(Address, Message)
}

func (e *EthService) EthSignTransaction(Transaction TransactionObjectWithSender) (result string, err error) {
	if e.EthSignTransactionFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthSignTransactionFunc(Transaction)
}

func (e *EthService) EthGetBalance(Address string, Block string) (result string, err error) {
	if e.EthGetBalanceFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetBalanceFunc(Address, Block)
}

func (e *EthService) EthGetStorage(Address string, StorageSlot string, Block string) (result string, err error) {
	if e.EthGetStorageFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetStorageFunc(Address, StorageSlot, Block)
}

func (e *EthService) EthGetTransactionCount(Address string, Block string) (result []string, err error) {
	if e.EthGetTransactionCountFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetTransactionCountFunc(Address, Block)
}

func (e *EthService) EthGetCode(Address string, Block string) (result string, err error) {
	if e.EthGetCodeFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetCodeFunc(Address, Block)
}

func (e *EthService) EthSendTransaction(Transaction TransactionObjectWithSender) (result string, err error) {
	if e.EthSendTransactionFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthSendTransactionFunc(Transaction)
}

func (e *EthService) EthSendRawTransaction(Transaction string) (result string, err error) {
	if e.EthSendRawTransactionFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthSendRawTransactionFunc(Transaction)
}

func (e *EthService) EthGetTransactionByHash(TransactionHash string) (result TransactionInformation, err error) {
	if e.EthGetTransactionByHashFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetTransactionByHashFunc(TransactionHash)
}

func (e *EthService) EthGetTransactionByBlockHashAndIndex(BlockHash string, TransactionIndex string) (result TransactionInformation, err error) {
	if e.EthGetTransactionByBlockHashAndIndexFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetTransactionByBlockHashAndIndexFunc(BlockHash, TransactionIndex)
}

func (e *EthService) EthGetTransactionByBlockNumberAndIndex(BlockNumber string, TransactionIndex string) (result TransactionInformation, err error) {
	if e.EthGetTransactionByBlockNumberAndIndexFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetTransactionByBlockNumberAndIndexFunc(BlockNumber, TransactionIndex)
}

func (e *EthService) EthGetTransactionReceipt(TransactionHash string) (result EthGetTransactionReceiptResult, err error) {
	if e.EthGetTransactionReceiptFunc == nil {
		return result, UnregisteredErr
	}
	return e.EthGetTransactionReceiptFunc(TransactionHash)
}
