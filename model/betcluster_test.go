package model

import (
	"MDBWeb/orm"
	"reflect"
	"testing"
)

func TestGetBetCluster(t *testing.T) {
	orm.OpenDB()
	orm.TableInit()
	type args struct {
		roundID string
	}
	tests := []struct {
		name string
		args args
		want *orm.BetCluster
	}{
		// TODO: Add test cases.
		{"case1", args{"qq"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBetCluster(tt.args.roundID); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBetCluster() = %v, want %v", got, tt.want)
			}
		})
	}
}
