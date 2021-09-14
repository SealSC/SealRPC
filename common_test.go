package SealRPC

import (
	"reflect"
	"testing"
)

func TestLogResult_MarshalJSON(t *testing.T) {
	type fields struct {
		Arr  []string
		Logs []Log
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{name: "empty", fields: fields{Arr: []string{}, Logs: []Log{}}, want: []byte("[]"), wantErr: false},
		{name: "stringArr", fields: fields{Arr: []string{"a", "b", "c"}, Logs: []Log{}}, want: []byte(`["a","b","c"]`), wantErr: false},
		{name: "StructArr", fields: fields{Arr: []string{}, Logs: []Log{{
			Topics:           []string{"0x1", "0x2"},
			LogIndex:         "1",
			Address:          "0x3",
			TransactionHash:  "0x123",
			BlockHash:        "0x123456",
			BlockNumber:      "0x11",
			Data:             "0x111",
			Removed:          true,
			TransactionIndex: "0x112233",
		}}}, want: []byte(`[{"topics":["0x1","0x2"],"logIndex":"1","address":"0x3","transactionHash":"0x123","blockHash":"0x123456","blockNumber":"0x11","data":"0x111","removed":true,"transactionIndex":"0x112233"}]`), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := LogResult{
				Arr:  tt.fields.Arr,
				Logs: tt.fields.Logs,
			}
			got, err := l.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestLogResult_UnmarshalJSON(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    LogResult
		wantErr bool
	}{
		{name: "empty", args: args{bytes: []byte("[]")}, want: LogResult{Arr: []string{}}, wantErr: false},
		{name: "stringArr", args: args{bytes: []byte(`["a","b","c"]`)}, want: LogResult{Arr: []string{"a", "b", "c"}, Logs: nil}, wantErr: false},
		{name: "StructArr", args: args{bytes: []byte(`[{"topics":["0x1","0x2"],"logIndex":"1","address":"0x3","transactionHash":"0x123","blockHash":"0x123456","blockNumber":"0x11","data":"0x111","removed":true,"transactionIndex":"0x112233"}]`)}, want: LogResult{Arr: nil, Logs: []Log{{
			Topics:           []string{"0x1", "0x2"},
			LogIndex:         "1",
			Address:          "0x3",
			TransactionHash:  "0x123",
			BlockHash:        "0x123456",
			BlockNumber:      "0x11",
			Data:             "0x111",
			Removed:          true,
			TransactionIndex: "0x112233",
		}}}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &LogResult{}
			if err := l.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(*l, tt.want) {
				t.Errorf("MarshalJSON() l = %v, want %v", *l, tt.want)
			}
		})
	}
}

func TestStringArr_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		s       StringArr
		args    args
		wantErr bool
	}{
		{name: "simpleString", s: nil, args: args{data: []byte(`"456"`)}, wantErr: false},
		{name: "stringArrays", s: nil, args: args{data: []byte(`["123","456"]`)}, wantErr: false},
		{name: "stringArray(len==1)", s: nil, args: args{data: []byte(`["123","456"]`)}, wantErr: false},
		{name: "jsonObject", s: nil, args: args{data: []byte(`{"name":"tom"}`)}, wantErr: true},
		{name: "not a json", s: nil, args: args{data: []byte(`xxx hhe!@#`)}, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSyncing_MarshalJSON(t *testing.T) {
	type fields struct {
		CurrentBlock  string
		HighestBlock  string
		StartingBlock string
		ISSyncing     bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{name: "Syncing", fields: fields{CurrentBlock: "0x111", HighestBlock: "0x222", StartingBlock: "0x333", ISSyncing: true}, want: []byte(`{"currentBlock":"0x111","highestBlock":"0x222","startingBlock":"0x333"}`), wantErr: false},
		{name: "NotSyncing", fields: fields{CurrentBlock: "0x111", HighestBlock: "0x123", StartingBlock: "0x12456", ISSyncing: false}, want: []byte(`false`), wantErr: false},
		{name: "Syncing but empty", fields: fields{CurrentBlock: "", HighestBlock: "", StartingBlock: "", ISSyncing: true}, want: []byte(`{"currentBlock":"","highestBlock":"","startingBlock":""}`), wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Syncing{
				Sync: sync{
					CurrentBlock:  tt.fields.CurrentBlock,
					HighestBlock:  tt.fields.HighestBlock,
					StartingBlock: tt.fields.StartingBlock,
				},
				ISSyncing: tt.fields.ISSyncing,
			}
			got, err := e.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}

func TestSyncing_UnmarshalJSON(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Syncing
		wantErr bool
	}{
		{name: "Syncing", args: args{bytes: []byte(`{"currentBlock":"0x111","highestBlock":"0x222","startingBlock":"0x333"}`)}, want: Syncing{Sync: sync{CurrentBlock: "0x111", HighestBlock: "0x222", StartingBlock: "0x333"}, ISSyncing: true}, wantErr: false},
		{name: "NotSyncing", args: args{[]byte(`false`)}, want: Syncing{ISSyncing: false}, wantErr: false},
		{name: "empty", args: args{[]byte(`{"currentBlock":"","highestBlock":"","startingBlock":""}`)}, want: Syncing{Sync: sync{CurrentBlock: "", HighestBlock: "", StartingBlock: ""}, ISSyncing: true}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Syncing{}
			if err := e.UnmarshalJSON(tt.args.bytes); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !reflect.DeepEqual(*e, tt.want) {
				t.Errorf("MarshalJSON() e = %v, want %v", *e, tt.want)
			}
		})
	}
}
