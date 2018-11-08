package model

import (
	"MDBWeb/orm"
	"testing"
)

func Test_getSumOfFeatureWin(t *testing.T) {
	orm.OpenDB()
	orm.TableInit()
	type args struct {
		clusterID   int64
		featureType int
		ps          int
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
		{"case1", args{4705, 3, 14}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSumOfFeatureWin(tt.args.clusterID, tt.args.featureType, tt.args.ps); got != tt.want {
				t.Errorf("getSumOfFeatureWin() = %v, want %v", got, tt.want)
			}
		})
	}
}
