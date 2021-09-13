package SealRPC

import (
	"bytes"
	"encoding/json"
)

type EthInterface interface {
	// EthGetBlockByHash Returns information about a block by hash.
	// BlockHash:32 byte hex value
	// HydratedTransactions:hydrated
	EthGetBlockByHash(BlockHash string, HydratedTransactions bool) (result Block, err error)
	// EthGetBlockByNumber Returns information about a block by number.
	// BlockNumber:hex encoded unsigned integer
	// HydratedTransactions:hydrated
	EthGetBlockByNumber(BlockNumber string, HydratedTransactions bool) (result Block, err error)
	// EthGetBlockTransactionCountByHash Returns the number of transactions in a block from a block matching the given block hash.
	// BlockHash:32 byte hex value
	// Return: Transaction count[hex encoded unsigned integer]
	EthGetBlockTransactionCountByHash(BlockHash string) (result []string, err error)
	// EthGetBlockTransactionCountByNumber Returns the number of transactions in a block matching the given block number.
	// BlockNumber:hex encoded unsigned integer
	// Return: Transaction count[hex encoded unsigned integer]
	EthGetBlockTransactionCountByNumber(BlockNumber string) (result []string, err error)
	// EthGetUncleCountByBlockHash Returns the number of uncles in a block from a block matching the given block hash.
	// BlockHash:32 byte hex value
	// Return: Uncle count[hex encoded unsigned integer]
	EthGetUncleCountByBlockHash(BlockHash string) (result []string, err error)
	// EthGetUncleCountByBlockNumber Returns the number of transactions in a block matching the given block number.
	// BlockNumber:hex encoded unsigned integer
	// Return: Uncle count[hex encoded unsigned integer]
	EthGetUncleCountByBlockNumber(BlockNumber string) (result []string, err error)
	// EthProtocolVersion Returns the current ethereum protocol version.
	// Return: version
	EthProtocolVersion() (result string, err error)
	// EthSyncing Returns an object with data about the sync status or false.
	EthSyncing() (result Syncing, err error)
	// EthCoinbase Returns the client coinbase address.
	// Return: hex encoded address
	EthCoinbase() (result string, err error)
	// EthAccounts Returns a list of addresses owned by client.
	// Returns: Accounts[hex encoded address]
	EthAccounts() (result []string, err error)
	// EthBlockNumber Returns the number of most recent block.
	// Return: hex encoded unsigned integer
	EthBlockNumber() (result string, err error)
	// EthCall Executes a new message call immediately without creating a transaction on the block chain.
	// Transaction:Transaction object with sender
	// Return: hex encoded bytes
	EthCall(Transaction TransactionObjectWithSender) (result string, err error)
	// EthEstimateGas Generates and returns an estimate of how much gas is necessary to allow the transaction to complete.
	// Transaction:Transaction object with sender
	// Return: hex encoded unsigned integer
	EthEstimateGas(Transaction TransactionObjectWithSender) (result string, err error)
	// EthGasPrice Returns the current price per gas in wei.
	// Return: Gas price
	EthGasPrice() (result string, err error)
	// EthFeeHistory
	//
	// BlockCount:hex encoded unsigned integer
	// NewestBlock:Block number or tag
	// RewardPercentiles:rewardPercentiles
	EthFeeHistory(BlockCount string, NewestBlock string, RewardPercentiles []int) (result EthFeeHistoryResult, err error)
	// EthNewFilter Creates a filter object, based on filter options, to notify when the state changes (logs).
	// Filter:filter
	// Return: hex encoded unsigned integer
	EthNewFilter(Filter Filter) (result string, err error)
	// EthNewBlockFilter Creates a filter in the node, to notify when a new block arrives.
	// Return: hex encoded unsigned integer
	EthNewBlockFilter() (result string, err error)
	// EthNewPendingTransactionFilter Creates a filter in the node, to notify when new pending transactions arrive.
	// Return: hex encoded unsigned integer
	EthNewPendingTransactionFilter() (result string, err error)
	// EthUninstallFilter Uninstalls a filter with given id.
	// FilterIdentifier:hex encoded unsigned integer
	EthUninstallFilter(FilterIdentifier string) (result bool, err error)
	// EthGetFilterChanges Polling method for a filter, which returns an array of logs which occurred since last poll.
	// FilterIdentifier:hex encoded unsigned integer
	EthGetFilterChanges(FilterIdentifier string) (result LogResult, err error)
	// EthGetFilterLogs Returns an array of all logs matching filter with given id.
	// FilterIdentifier:hex encoded unsigned integer
	EthGetFilterLogs(FilterIdentifier string) (result LogResult, err error)
	// EthGetLogs Returns an array of all logs matching filter with given id.
	// Filter:filter
	EthGetLogs(Filter Filter) (result LogResult, err error)
	// EthMining Returns whether the client is actively mining new blocks.
	// Return: miningStatus
	EthMining() (result bool, err error)
	// EthHashrate Returns the number of hashes per second that the node is mining with.
	// Return: Hashrate
	EthHashrate() (result string, err error)
	// EthGetWork Returns the hash of the current block, the seedHash, and the boundary condition to be met (“target”).
	// Return: Proof-of-work hash
	EthGetWork() (result []string, err error)
	// EthSubmitWork Used for submitting a proof-of-work solution.
	// ProofOfWorkHash:32 hex encoded bytes
	// SeedHash:32 hex encoded bytes
	// Difficulty:32 hex encoded bytes
	EthSubmitWork(ProofOfWorkHash string, SeedHash string, Difficulty string) (result bool, err error)
	// EthSubmitHashrate Used for submitting mining hashrate.
	// Hashrate:32 hex encoded bytes
	// ID:32 hex encoded bytes
	EthSubmitHashrate(Hashrate string, ID string) (result bool, err error)
	// EthSign Returns an EIP-191 signature over the provided data.
	// Address:hex encoded address
	// Message:hex encoded bytes
	// Return: 65 hex encoded bytes
	EthSign(Address string, Message string) (result string, err error)
	// EthSignTransaction Returns an RLP encoded transaction signed by the specified account.
	// Transaction:Transaction object with sender
	// Return: hex encoded bytes
	EthSignTransaction(Transaction TransactionObjectWithSender) (result string, err error)
	// EthGetBalance Returns the balance of the account of given address.
	// Address:hex encoded address
	// Block:Block number or tag
	// Return: hex encoded unsigned integer
	EthGetBalance(Address string, Block string) (result string, err error)
	// EthGetStorage Returns the value from a storage position at a given address.
	// Address:hex encoded address
	// StorageSlot:hex encoded unsigned integer
	// Block:Block number or tag
	// Return: hex encoded bytes
	EthGetStorage(Address string, StorageSlot string, Block string) (result string, err error)
	// EthGetTransactionCount Returns the number of transactions sent from an address.
	// Address:hex encoded address
	// Block:Block number or tag
	// Return: Transaction count[hex encoded unsigned integer]
	EthGetTransactionCount(Address string, Block string) (result []string, err error)
	// EthGetCode Returns code at a given address.
	// Address:hex encoded address
	// Block:Block number or tag
	// Return: hex encoded bytes
	EthGetCode(Address string, Block string) (result string, err error)
	// EthSendTransaction Signs and submits a transaction.
	// Transaction:Transaction object with sender
	// Return: 32 byte hex value
	EthSendTransaction(Transaction TransactionObjectWithSender) (result string, err error)
	// EthSendRawTransaction Submits a raw transaction.
	// Transaction:hex encoded bytes
	// Return: 32 byte hex value
	EthSendRawTransaction(Transaction string) (result string, err error)
	// EthGetTransactionByHash Returns the information about a transaction requested by transaction hash.
	// TransactionHash:32 byte hex value
	// Return: Transaction information
	EthGetTransactionByHash(TransactionHash string) (result TransactionInformation, err error)
	// EthGetTransactionByBlockHashAndIndex Returns information about a transaction by block hash and transaction index position.
	// BlockHash:32 byte hex value
	// TransactionIndex:hex encoded unsigned integer
	// Return: Transaction information
	EthGetTransactionByBlockHashAndIndex(BlockHash string, TransactionIndex string) (result TransactionInformation, err error)
	// EthGetTransactionByBlockNumberAndIndex Returns information about a transaction by block number and transaction index position.
	// BlockNumber:hex encoded unsigned integer
	// TransactionIndex:hex encoded unsigned integer
	// Return: Transaction information
	EthGetTransactionByBlockNumberAndIndex(BlockNumber string, TransactionIndex string) (result TransactionInformation, err error)
	// EthGetTransactionReceipt Returns the receipt of a transaction by transaction hash.
	// TransactionHash:32 byte hex value
	EthGetTransactionReceipt(TransactionHash string) (result EthGetTransactionReceiptResult, err error)
}

type StringArr []string

func (s *StringArr) UnmarshalJSON(data []byte) error {
	data = bytes.TrimSpace(data)
	if len(data) <= 0 {
		return nil
	}
	if data[0] == '"' {
		var temp string
		if err := json.Unmarshal(data, &temp); err != nil {
			return err
		}
		*s = []string{temp}
		return nil
	} else {
		return json.Unmarshal(data, s)
	}
}

type TransactionAccess struct {
	Address     string   `json:"address"`     //  hex encoded address
	StorageKeys []string `json:"storageKeys"` //  [32 byte hex value]
}
type TransactionHexByte string
type Transaction1559 struct {
	From                 string              `json:"from"`                 //  文档丢失了这个字段
	Type                 string              `json:"type"`                 //  type
	Nonce                string              `json:"nonce"`                //  nonce
	To                   string              `json:"to"`                   //  to address
	Gas                  string              `json:"gas"`                  //  gas limit
	Value                string              `json:"value"`                //  value
	Input                string              `json:"input"`                //  input data
	MaxPriorityFeePerGas string              `json:"maxPriorityFeePerGas"` // Maximum fee per gas the sender is willing to pay to miners in wei max priority fee per gas
	MaxFeePerGas         string              `json:"maxFeePerGas"`         // The maximum total fee per gas the sender is willing to pay (includes the network / base fee and miner / priority fee) in wei max fee per gas
	AccessList           []TransactionAccess `json:"accessList"`           // EIP-2930 access list accessList[Access list entry]
	ChainId              string              `json:"chainId"`              // Chain ID that this transaction is valid on. chainId
	YParity              string              `json:"yParity"`              // The parity (0 for even, 1 for odd) of the y-value of the secp256k1 signature. yParity
	R                    string              `json:"r"`                    //  r
	S                    string              `json:"s"`                    //  s
}
type Transaction2930 struct {
	From       string              `json:"from"`       //  文档丢失了这个字段
	Type       string              `json:"type"`       //  type
	Nonce      string              `json:"nonce"`      //  nonce
	To         string              `json:"to"`         //  to address
	Gas        string              `json:"gas"`        //  gas limit
	Value      string              `json:"value"`      //  value
	Input      string              `json:"input"`      //  input data
	GasPrice   string              `json:"gasPrice"`   // The gas price willing to be paid by the sender in wei gas price
	AccessList []TransactionAccess `json:"accessList"` // EIP-2930 access list accessList[Access list entry]
	ChainId    string              `json:"chainId"`    // Chain ID that this transaction is valid on. chainId
	YParity    string              `json:"yParity"`    // The parity (0 for even, 1 for odd) of the y-value of the secp256k1 signature. yParity
	R          string              `json:"r"`          //  r
	S          string              `json:"s"`          //  s
}
type TransactionLegacy struct {
	From     string `json:"from"`     //  文档丢失了这个字段
	Type     string `json:"type"`     //  type
	Nonce    string `json:"nonce"`    //  nonce
	To       string `json:"to"`       //  to address
	Gas      string `json:"gas"`      //  gas limit
	Value    string `json:"value"`    //  value
	Input    string `json:"input"`    //  input data
	GasPrice string `json:"gasPrice"` // The gas price willing to be paid by the sender in wei gas price
	ChainId  string `json:"chainId"`  // Chain ID that this transaction is valid on. chainId
	V        string `json:"v"`        //  v
	R        string `json:"r"`        //  r
	S        string `json:"s"`        //  s
}

type Transactions interface {
	//TODO 需要适配Transaction多种类型
}

type Syncing struct {
	CurrentBlock  string `json:"currentBlock"`  //  Current block
	HighestBlock  string `json:"highestBlock"`  //  Highest block
	StartingBlock string `json:"startingBlock"` //  Starting block
	ISSyncing     bool   `json:"is_syncing"`
}

func (e Syncing) MarshalJSON() ([]byte, error) {
	if !e.ISSyncing {
		return json.Marshal(e.ISSyncing)
	}
	return json.Marshal(e)
}

func (e *Syncing) UnmarshalJSON(bytes []byte) error {
	if err := json.Unmarshal(bytes, &e); err != nil {
		return json.Unmarshal(bytes, &e.ISSyncing)
	}
	return nil
}

// EthFeeHistoryResult Fee history results. feeHistoryResults
type EthFeeHistoryResult struct {
	OldestBlock   string     `json:"oldestBlock"`   // Lowest number block of returned range. oldestBlock
	BaseFeePerGas []string   `json:"baseFeePerGas"` // An array of block base fees per gas. This includes the next block after the newest of the returned range, because this value can be derived from the newest block. Zeroes are returned for pre-EIP-1559 blocks. baseFeePerGasArray[hex encoded unsigned integer]
	Reward        [][]string `json:"reward"`        // A two-dimensional array of effective priority fees per gas at the requested block percentiles. rewardArray[rewardPercentile]
}

// EthGetTransactionReceiptResult Filter results [Receipt info]
type EthGetTransactionReceiptResult struct {
	TransactionIndex  string `json:"transactionIndex"`  //  transaction index
	BlockHash         string `json:"blockHash"`         //  block hash
	From              string `json:"from"`              //  from
	To                string `json:"to"`                //  to
	LogsBloom         string `json:"logsBloom"`         //  logs bloom
	TransactionHash   string `json:"transactionHash"`   //  transaction hash
	BlockNumber       string `json:"blockNumber"`       //  block number
	CumulativeGasUsed string `json:"cumulativeGasUsed"` //  accessList
	GasUsed           string `json:"gasUsed"`           //  gasUsed
	ContractAddress   string `json:"contractAddress"`   //  contract address
	Logs              []Log  `json:"logs"`              //  logs[log]
}

// Block object
type Block struct {
	Number           *string      `json:"number"`           //  Number
	Hash             *string      `json:"hash"`             //  keccak
	ParentHash       *string      `json:"parentHash"`       //  Parent block hash
	Nonce            *string      `json:"nonce"`            //  Nonce
	Sha3Uncles       string       `json:"sha3Uncles"`       //  Ommers hash
	LogsBloom        string       `json:"logsBloom"`        //  Bloom filter
	TransactionsRoot string       `json:"transactionsRoot"` //  Transactions root
	StateRoot        string       `json:"stateRoot"`        //  State root
	Miner            *string      `json:"miner"`            //  Coinbase
	Difficulty       string       `json:"difficulty"`       //  Difficulty
	TotalDifficulty  *string      `json:"totalDifficulty"`  //  Total difficult
	ExtraData        string       `json:"extraData"`        //  Extra data
	Size             string       `json:"size"`             //  Block size
	GasLimit         string       `json:"gasLimit"`         //  Gas limit
	GasUsed          string       `json:"gasUsed"`          //  Gas used
	Timestamp        string       `json:"timestamp"`        //  Timestamp
	Transactions     Transactions `json:"transactions"`     //
	Uncles           []string     `json:"uncles"`           //  Uncles[32 byte hex value]
	BaseFeePerGas    string       `json:"baseFeePerGas"`    //  Base fee per gas
	ReceiptsRoot     string       `json:"receiptsRoot"`     //  Receipts root
	MixHash          string       `json:"mixHash"`          //  Mix hash
}

// TransactionObjectWithSender Transaction object with sender
type TransactionObjectWithSender struct {
	From string `json:"from"` //  from
	Transactions
}

// Filter
type Filter struct {
	FromBlock string    `json:"fromBlock"` //  from block
	ToBlock   string    `json:"toBlock"`   //  to block
	Address   StringArr `json:"address"`   //  Address(es) Address|Addresses[hex encoded address]
	Topics    []string  `json:"topics"`    //  Topics[Topic]
}
type Log struct {
	Topics           []string `json:"topics"`           //  topics[Topic]
	LogIndex         string   `json:"logIndex"`         //  log index
	Address          string   `json:"address"`          //  address
	TransactionHash  string   `json:"transactionHash"`  //  transaction hash
	BlockHash        string   `json:"blockHash"`        //  block hash
	BlockNumber      string   `json:"blockNumber"`      //  block number
	Data             string   `json:"data"`             //  data
	Removed          bool     `json:"removed"`          //  removed
	TransactionIndex string   `json:"transactionIndex"` //  transaction index
}

// LogResult Filter results
type LogResult struct {
	Arr  []string // new block hashes[32 byte hex value]  ||  new transaction hashes[32 byte hex value]
	Logs []Log    // new logs[log]
}

func (l *LogResult) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &l.Arr)
	if err != nil {
		return json.Unmarshal(bytes, &l.Logs)
	}
	return nil
}

func (l LogResult) MarshalJSON() ([]byte, error) {
	if len(l.Arr) == 0 {
		return json.Marshal(l.Logs)
	}
	return json.Marshal(l.Arr)
}

// TransactionInformation Transaction information
type TransactionInformation struct {
	Hash             string `json:"hash"`             //  transaction hash
	TransactionIndex string `json:"transactionIndex"` //  transaction index
	BlockHash        string `json:"blockHash"`        //  block hash
	BlockNumber      string `json:"blockNumber"`      //  block number
	From             string `json:"from"`             //  from address
}
