package SealRPC

import (
	"bytes"
	"encoding/json"
)

type EthInterface interface {
	EthGetBlockByHash(BlockHash string, HydratedTransactions bool) (result Block, err error)
	EthGetBlockByNumber(BlockNumber string, HydratedTransactions bool) (result Block, err error)
	EthGetBlockTransactionCountByHash(BlockHash string) (result []string, err error)
	EthGetBlockTransactionCountByNumber(BlockNumber string) (result []string, err error)
	EthGetUncleCountByBlockHash(BlockHash string) (result []string, err error)
	EthGetUncleCountByBlockNumber(BlockNumber string) (result []string, err error)
	EthProtocolVersion() (result string, err error)
	EthSyncing() (result Syncing, err error)
	EthCoinbase() (result string, err error)
	EthAccounts() (result []string, err error)
	EthBlockNumber() (result string, err error)
	EthCall(Transaction TransactionObjectWithSender) (result string, err error)
	EthEstimateGas(Transaction TransactionObjectWithSender) (result string, err error)
	EthGasPrice() (result string, err error)
	EthFeeHistory(BlockCount string, NewestBlock string, RewardPercentiles []int) (result EthFeeHistoryResult, err error)
	EthNewFilter(Filter Filter) (result string, err error)
	EthNewBlockFilter() (result string, err error)
	EthNewPendingTransactionFilter() (result string, err error)
	EthUninstallFilter(FilterIdentifier string) (result bool, err error)
	EthGetFilterChanges(FilterIdentifier string) (result LogResult, err error)
	EthGetFilterLogs(FilterIdentifier string) (result LogResult, err error)
	EthGetLogs(Filter Filter) (result LogResult, err error)
	EthMining() (result bool, err error)
	EthHashrate() (result string, err error)
	EthGetWork() (result []string, err error)
	EthSubmitWork(ProofOfWorkHash string, SeedHash string, Difficulty string) (result bool, err error)
	EthSubmitHashrate(Hashrate string, ID string) (result bool, err error)
	EthSign(Address string, Message string) (result string, err error)
	EthSignTransaction(Transaction TransactionObjectWithSender) (result string, err error)
	EthGetBalance(Address string, Block string) (result string, err error)
	EthGetStorage(Address string, StorageSlot string, Block string) (result string, err error)
	EthGetTransactionCount(Address string, Block string) (result []string, err error)
	EthGetCode(Address string, Block string) (result string, err error)
	EthSendTransaction(Transaction TransactionObjectWithSender) (result string, err error)
	EthSendRawTransaction(Transaction string) (result string, err error)
	EthGetTransactionByHash(TransactionHash string) (result TransactionInformation, err error)
	EthGetTransactionByBlockHashAndIndex(BlockHash string, TransactionIndex string) (result TransactionInformation, err error)
	EthGetTransactionByBlockNumberAndIndex(BlockNumber string, TransactionIndex string) (result TransactionInformation, err error)
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

type EthFeeHistoryResult struct {
	OldestBlock   string     `json:"oldestBlock"`   // Lowest number block of returned range. oldestBlock
	BaseFeePerGas []string   `json:"baseFeePerGas"` // An array of block base fees per gas. This includes the next block after the newest of the returned range, because this value can be derived from the newest block. Zeroes are returned for pre-EIP-1559 blocks. baseFeePerGasArray[hex encoded unsigned integer]
	Reward        [][]string `json:"reward"`        // A two-dimensional array of effective priority fees per gas at the requested block percentiles. rewardArray[rewardPercentile]
} // Fee history results. feeHistoryResults

//  Filter results
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
} //  Receipt info
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
} //  Block object
type TransactionObjectWithSender struct {
	From string `json:"from"` //  from
	Transactions
} //  Transaction object with sender

type Filter struct {
	FromBlock string    `json:"fromBlock"` //  from block
	ToBlock   string    `json:"toBlock"`   //  to block
	Address   StringArr `json:"address"`   //  Address(es) Address|Addresses[hex encoded address]
	Topics    []string  `json:"topics"`    //  Topics[Topic]
} //  filter
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
type LogResult struct {
	Arr  []string // new block hashes[32 byte hex value]  ||  new transaction hashes[32 byte hex value]
	Logs []Log    // new logs[log]
} //  Filter results

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

type TransactionInformation struct {
	Hash             string `json:"hash"`             //  transaction hash
	TransactionIndex string `json:"transactionIndex"` //  transaction index
	BlockHash        string `json:"blockHash"`        //  block hash
	BlockNumber      string `json:"blockNumber"`      //  block number
	From             string `json:"from"`             //  from address
} //  Transaction information
