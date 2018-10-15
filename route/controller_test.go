package route

import "testing"

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
