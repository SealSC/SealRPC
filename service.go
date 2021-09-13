package SealRPC

import (
	"errors"
)

var UnregisteredErr = errors.New("function Unregistered")

type EthService struct {
	//as EthInterface.EthGetBlockByHas
	EthGetBlockByHashFuc func(BlockHash string, HydratedTransactions bool) (Block, error)
	//as EthInterface.EthGetBlockByNumber
	EthGetBlockByNumberFuc func(BlockNumber string, HydratedTransactions bool) (Block, error)
	//as EthInterface.EthGetBlockTransactionCountByHas
	EthGetBlockTransactionCountByHashFuc func(BlockHash string) ([]string, error)
	//as EthInterface.EthGetBlockTransactionCountByNumber
	EthGetBlockTransactionCountByNumberFunc func(BlockNumber string) ([]string, error)
	//as EthInterface.EthGetUncleCountByBlockHash
	EthGetUncleCountByBlockHashFunc func(BlockHash string) ([]string, error)
	//as EthInterface.EthGetUncleCountByBlockNumber
	EthGetUncleCountByBlockNumberFunc func(BlockNumber string) ([]string, error)
	//as EthInterface.EthProtocolVersion
	EthProtocolVersionFunc func() (string, error)
	//as EthInterface.EthSyncing
	EthSyncingFunc func() (Syncing, error)
	//as EthInterface.EthCoinbase
	EthCoinbaseFunc func() (string, error)
	//as EthInterface.EthAccounts
	EthAccountsFunc func() ([]string, error)
	//as EthInterface.EthBlockNumber
	EthBlockNumberFunc func() (string, error)
	//as EthInterface.EthCall
	EthCallFunc func(Transaction TransactionObjectWithSender) (string, error)
	//as EthInterface.EthEstimateGas
	EthEstimateGasFunc func(Transaction TransactionObjectWithSender) (string, error)
	//as EthInterface.EthGasPrice
	EthGasPriceFunc func() (string, error)
	//as EthInterface.EthFeeHistory
	EthFeeHistoryFunc func(BlockCount string, NewestBlock string, RewardPercentiles []int) (EthFeeHistoryResult, error)
	//as EthInterface.EthNewFilter
	EthNewFilterFunc func(Filter Filter) (string, error)
	//as EthInterface.EthNewBlockFilter
	EthNewBlockFilterFunc func() (string, error)
	//as EthInterface.EthNewPendingTransactionFilter
	EthNewPendingTransactionFilterFunc func() (string, error)
	//as EthInterface.EthUninstallFilter
	EthUninstallFilterFunc func(FilterIdentifier string) (bool, error)
	//as EthInterface.EthGetFilterChanges
	EthGetFilterChangesFunc func(FilterIdentifier string) (LogResult, error)
	//as EthInterface.EthGetFilterLogs
	EthGetFilterLogsFunc func(FilterIdentifier string) (LogResult, error)
	//as EthInterface.EthGetLogs
	EthGetLogsFunc func(Filter Filter) (LogResult, error)
	//as EthInterface.EthMining
	EthMiningFunc func() (bool, error)
	//as EthInterface.EthHashrate
	EthHashrateFunc func() (string, error)
	//as EthInterface.EthGetWork
	EthGetWorkFunc func() ([]string, error)
	//as EthInterface.EthSubmitWork
	EthSubmitWorkFunc func(ProofOfWorkHash string, SeedHash string, Difficulty string) (bool, error)
	//as EthInterface.EthSubmitHashrate
	EthSubmitHashrateFunc func(Hashrate string, ID string) (bool, error)
	//as EthInterface.EthSign
	EthSignFunc func(Address string, Message string) (string, error)
	//as EthInterface.EthSignTransaction
	EthSignTransactionFunc func(Transaction TransactionObjectWithSender) (string, error)
	//as EthInterface.EthGetBalance
	EthGetBalanceFunc func(Address string, Block string) (string, error)
	//as EthInterface.EthGetStorage
	EthGetStorageFunc func(Address string, StorageSlot string, Block string) (string, error)
	//as EthInterface.EthGetTransactionCount
	EthGetTransactionCountFunc func(Address string, Block string) ([]string, error)
	//as EthInterface.EthGetCode
	EthGetCodeFunc func(Address string, Block string) (string, error)
	//as EthInterface.EthSendTransaction
	EthSendTransactionFunc func(Transaction TransactionObjectWithSender) (string, error)
	//as EthInterface.EthSendRawTransaction
	EthSendRawTransactionFunc func(Transaction string) (string, error)
	//as EthInterface.EthGetTransactionByHash
	EthGetTransactionByHashFunc func(TransactionHash string) (TransactionInformation, error)
	//as EthInterface.EthGetTransactionByBlockHashAndIndex
	EthGetTransactionByBlockHashAndIndexFunc func(BlockHash string, TransactionIndex string) (TransactionInformation, error)
	//as EthInterface.EthGetTransactionByBlockNumberAndIndex
	EthGetTransactionByBlockNumberAndIndexFunc func(BlockNumber string, TransactionIndex string) (TransactionInformation, error)
	//as EthInterface.EthGetTransactionReceipt
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
