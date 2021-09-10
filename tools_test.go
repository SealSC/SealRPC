package SealRPC

import "testing"

func TestCamelCase(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"0", args{"xx"}, "Xx"},
		{"1", args{"  xx  "}, "Xx"},
		{"2", args{"x  x"}, "XX"},
		{"3", args{"x---x"}, "XX"},
		{"4", args{"x__-x"}, "XX"},
		{"5", args{"x  _-x"}, "XX"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCase(tt.args.s); got != tt.want {
				t.Errorf("CamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
