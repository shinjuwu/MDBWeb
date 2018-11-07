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
		TotalFeatureHit: 0,
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
		TotalFeatureHit: 0,
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
		Result:          "eJy81MFqhDAQxvF3+c45jO4mu+ZVShHBiKFpLCZSivjuBauUabUHa/aYMPCDf8KM6M1rZX3Z1XWAJoHGVHHoTWlraCgIvFjnSuubDvppRIhVHAJ0JtDY0M5jmVpP8ePNQGc3gXfroSURkUBojXPzhR+c+ya+pkH4pU5iR7rfuSQX6XK6VBCT8gW6nuAQc360W6EDCv2h5Nvd/q3wZtcHNZPMuaRJphii0iA3/i5pkIIhRQokp50dcLIiN3/YMeR5+gwAAP//aOia4Q==",
		TotalFeatureHit: 0,
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
	ProcessCQ9Log()
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
