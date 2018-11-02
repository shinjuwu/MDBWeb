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
		Bet:             1000,
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
		Bet:             1000,
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
		Bet:             1000,
		FeatureType:     1,
		FishType:        "1",
		Result:          "eJy01UFqwzAQBdC7/LUWM7EVRbpKKcZgmYqmTokUSgm5e6H1InJVKNFoaWyexfw/9hVn/zaGZThNU4QjhdmP6XL2Q5jg0EHhNRyPQ1jmE9zTFTGN6RLhWGEO8eX7sX63XqTPdw/HRuEjLHCGiO7En7sg/HrJTZVhbcuwroX3Jof1Cne1sOnK8EMnpnv4kMP9Cv8DpT/RQ45aCdPmZdiLmPlUScTUmbkTMTe1EjHzjDoR08rnzkTyITGx/ESZNksqkj1T32BDmXSLqZoWqJUvKjOV/wCVaotWcYMPH7OuOOnz7SsAAP//nm16MQ==",
		TotalFeatureBet: 0,
		TotalRound:      0,
		TotalBet:        0,
		TotalWin:        0,
		FishID:          "0",
	}
	featureLogs := make(map[int]map[int]orm.PreprocessLog)
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

func TestProcessCQ9Log(t *testing.T) {
	orm.OpenDB()
	orm.TableInit()
	err := ProcessCQ9Log()
	if err != nil {

	}
}

func Test_setProcessed(t *testing.T) {
	orm.OpenDB()
	orm.TableInit()
	log := &orm.BetCluster{
		ClusterID: 111,
	}
	type args struct {
		log *orm.BetCluster
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"case1", args{log}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setProcessed(tt.args.log)
		})
	}
}
