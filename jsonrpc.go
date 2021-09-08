package SealRPC

import (
	"encoding/json"
	"io"
	"net/http"
	"reflect"
	"strconv"
)

var (
	ParseErr          = RPCError{Code: -32700, Message: "Parse error"}
	InvalidRequestErr = RPCError{Code: -32600, Message: "Invalid Request"}
	MethodNotFoundErr = RPCError{Code: -32601, Message: "Method not found"}
	InvalidParamsErr  = RPCError{Code: -32602, Message: "Invalid params"}
	InternalErrorErr  = RPCError{Code: -32603, Message: "Internal error"}
	ServerErrorErr    = RPCError{Code: -3200, Message: "Server error"}
)

type JsonRPCService struct {
	ser  *EthService
	serV reflect.Value
}

func NewJsonRPCService(ser *EthService) *JsonRPCService {
	return &JsonRPCService{ser: ser, serV: reflect.ValueOf(ser)}
}

func (j *JsonRPCService) Run(addr string) error {
	return http.ListenAndServe(addr, j)
}
func (j *JsonRPCService) ReturnErr(e RPCError, w io.Writer) {
	response := RPCResponse{
		Error: &e,
	}
	_ = json.NewEncoder(w).Encode(response)
}

func (j *JsonRPCService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		request  RPCRequest
		response RPCResponse
	)
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		j.ReturnErr(ParseErr, w)
		return
	}
	if request.Invalid() {
		j.ReturnErr(InvalidRequestErr, w)
		return
	}

	up := toUp(request.Method)
	method := j.serV.MethodByName(up)
	if !method.IsValid() {
		j.ReturnErr(MethodNotFoundErr, w)
		return
	}
	if method.Type().NumIn() != len(request.Params) {
		j.ReturnErr(InvalidParamsErr, w)
		return
	}
	var params []reflect.Value
	for i := 0; i < method.Type().NumIn(); i++ {
		value := reflect.New(method.Type().In(i))
		params = append(params, value.Elem())
		if err := json.Unmarshal(request.Params[i], value.Interface()); err != nil {
			j.ReturnErr(ParseErr, w)
			return
		}
	}
	returns := method.Call(params)
	switch len(returns) {
	case 1:
		if err, ok := returns[0].Interface().(error); ok && err != nil {
			response.Error = &RPCError{Code: -3200, Message: err.Error()}
		}
	case 2:
		response.Result = returns[0].Interface()
		if err, ok := returns[1].Interface().(error); ok && err != nil {
			response.Error = &RPCError{Code: -3200, Message: err.Error()}
		}
	}
	_ = json.NewEncoder(w).Encode(response)
	w.WriteHeader(http.StatusOK)
}

type RPCRequest struct {
	Method  string            `json:"method"`
	Params  []json.RawMessage `json:"params,omitempty"`
	ID      int               `json:"id"`
	JSONRPC string            `json:"jsonrpc"`
}

func (R RPCRequest) Invalid() bool {
	if R.JSONRPC != "2.0" {
		return true
	}
	if R.Method == "" {
		return true
	}
	return false
}

type RPCResponse struct {
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result,omitempty"`
	Error   *RPCError   `json:"error,omitempty"`
	ID      int         `json:"id"`
}

type RPCError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// Error function is provided to be used as error object.
func (e *RPCError) Error() string {
	return strconv.Itoa(e.Code) + ":" + e.Message
}
