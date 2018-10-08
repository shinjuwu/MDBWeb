package preprocess

import (
	"MDBWeb/orm"
	"testing"
)

func Test_createPreprocessLog(t *testing.T) {
	orm.OpenDB()
	orm.TableInit()
	data := &orm.BetCluster{
		ClusterID: 1,
	}
	type args struct {
		data *orm.BetCluster
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"case1", args{data}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createPreprocessLog(tt.args.data)
		})
	}
}
