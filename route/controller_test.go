package route

import (
	"reflect"
	"testing"
)

func Test_getDetailToken(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"case1", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := getDetailToken(); (err != nil) != tt.wantErr {
				t.Errorf("getDetailToken() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getDetailOrderInfo(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want *ResDetailOrder
	}{
		// TODO: Add test cases.
		{"case1", args{""}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDetailOrderInfo(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDetailOrderInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
