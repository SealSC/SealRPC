package SealRPC

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestJsonRPCService_hitPath(t *testing.T) {
	type args struct {
		Path    string
		reqPath string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"0", args{"/", ""}, true},
		{"0", args{"", "/"}, true},
		{"0", args{"/s", "s"}, true},
		{"0", args{"s", "/s"}, true},
		{"0", args{"/s", "/  s"}, false},
		{"0", args{"/  s", "/%20%20%20s"}, false},
		{"0", args{"/  s", "/%20%20s"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewJsonRPCService(tt.args.Path, nil).hitPath(tt.args.reqPath); got != tt.want {
				t.Errorf("JsonRPCService.hitPath(%v,%v) = %v, want %v", tt.args.Path, tt.args.reqPath, got, tt.want)
			}
		})
	}
}

type TestServer struct{}

func (t *TestServer) MethodParameterNil() bool {
	return true
}
func (t *TestServer) MethodParameterType(q int, w int32, e string, r string) string {
	return fmt.Sprintf("%d%d%s%s", q, w, e, r)
}
func (t *TestServer) MethodReturnNil() interface{} { return nil }
func (t *TestServer) MethodNotFound() bool         { return true }

func (t *TestServer) MethodNoReturn()                          {}
func (t *TestServer) MethodSimpleReturnErr() error             { return errors.New("simple err") }
func (t *TestServer) MethodMultipleReturn() (bool, int, error) { return true, 99, nil }

func TestJsonRPCService_Run(t *testing.T) {
	testServer := TestServer{}
	handler := NewJsonRPCService("/rpc", &testServer)
	server := httptest.NewServer(handler)
	client := server.Client()
	var JsonRpcClientDo = func(req RPCRequest) (resp RPCResponse, err error) {
		var (
			body     []byte
			httpResp *http.Response
		)
		if body, err = json.Marshal(req); err != nil {
			err = fmt.Errorf("json marshal err:%v", err)
			return
		}
		if httpResp, err = client.Post(server.URL+"/"+handler.path, "application/json", bytes.NewBuffer(body)); err != nil {
			err = fmt.Errorf("client post request err:%v", err)
			return
		}
		if err = json.NewDecoder(httpResp.Body).Decode(&resp); err != nil {
			err = fmt.Errorf("read response body json decode err:%v", err)
			return
		}
		return
	}

	type args struct {
		req RPCRequest
	}
	tests := []struct {
		name string
		args args
		want RPCResponse
	}{
		{"JsonRPCNil", args{RPCRequest{JSONRPC: "", Method: "method_parameter_nil"}}, RPCResponse{Error: InvalidRequestErr}},
		{"MethodParameterNil", args{RPCRequest{JSONRPC: "2.0", Method: "method_parameter_nil"}}, RPCResponse{Result: true}},
		{"MethodParameterType", args{RPCRequest{JSONRPC: "2.0", Method: "method_parameter_type", Params: []json.RawMessage{[]byte("9"), []byte("8"), []byte("\"tom\""), []byte("\"jump\"")}}}, RPCResponse{Result: "98tomjump"}},
		{"MethodReturnNil", args{RPCRequest{JSONRPC: "2.0", Method: "method_return_nil"}}, RPCResponse{Result: nil}},
		{"MethodNotFound", args{RPCRequest{JSONRPC: "2.0", Method: "hello_method"}}, RPCResponse{Error: MethodNotFoundErr}},
		{"MethodName-Camel-Case", args{RPCRequest{JSONRPC: "2.0", Method: "MethodParameterNil"}}, RPCResponse{Result: true}},
		{"MethodNoReturn", args{RPCRequest{JSONRPC: "2.0", Method: "method_no_return"}}, RPCResponse{}},
		{"MethodMultipleReturn", args{RPCRequest{JSONRPC: "2.0", Method: "method_multiple_return"}}, RPCResponse{Error: MultipleReturnErr}},
		{"MethodSimpleReturnErr", args{RPCRequest{JSONRPC: "2.0", Method: "method_simple_return_err"}}, RPCResponse{Error: RPCError{Code: CommonJSONRpcErrCode, Message: "simple err"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := JsonRpcClientDo(tt.args.req)
			if err != nil || (resp.Result != tt.want.Result || resp.Error != tt.want.Error) {
				t.Errorf("client.Post(req = %v) (result= %v,error= %v) want %v", tt.args.req, resp.Result, resp.Error, tt.want)
			}
		})
	}
}
