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

func Test_insertFeatureLog(t *testing.T) {
	orm.OpenDB()
	orm.TableInit()
	data := orm.PreprocessLog{
		ClusterID:       11,
		RoundID:         "AP1234545",
		WinOdds:         1000,
		FeatureType:     1,
		FishType:        "1",
		Result:          "sdfdfhfghf",
		TotalFeatureBet: 0,
		TotalRound:      0,
		TotalBet:        0,
		TotalWin:        0,
		FishID:          "0",
	}
	data2 := orm.PreprocessLog{
		ClusterID:       11,
		RoundID:         "AP1234545",
		WinOdds:         1000,
		FeatureType:     1,
		FishType:        "1",
		Result:          "sdfdfhfghf",
		TotalFeatureBet: 0,
		TotalRound:      0,
		TotalBet:        0,
		TotalWin:        0,
		FishID:          "0",
	}
	logs := map[int]orm.PreprocessLog{
		1: data,
		2: data2,
	}
	featureLogs := map[int]map[int]orm.PreprocessLog{
		100: logs,
	}
	type args struct {
		featureLogs map[int]map[int]orm.PreprocessLog
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"case1", args{featureLogs}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			insertFeatureLog(tt.args.featureLogs)
		})
	}
}

func Test_processFeatureLog(t *testing.T) {
	log := &orm.PreprocessLog{
		ClusterID:       11,
		RoundID:         "AP1234545",
		WinOdds:         1000,
		FeatureType:     1,
		FishType:        "1",
		Result:          "eJxEjU0KAjEMhe/y1l2kLnMVkWGKKQZDFZsiUnp3oeh09/74XkcSB0ciCkjNTHzzz1PAy+sVjIiAu5ptWvIDfO6ovnur4BiQtd7m7K8PxFsL+DTpWXZvL/l1IKxoPhDGZXwDAAD//+4TLZ4=",
		TotalFeatureBet: 0,
		TotalRound:      0,
		TotalBet:        0,
		TotalWin:        0,
		FishID:          "0",
	}
	featureLogs := map[int]map[int]orm.PreprocessLog{}
	type args struct {
		log         *orm.PreprocessLog
		featureLogs map[int]map[int]orm.PreprocessLog
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"case1", args{log, featureLogs}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			processFeatureLog(tt.args.log, tt.args.featureLogs)
		})
	}
}
