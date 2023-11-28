package cmd

import "testing"

func TestAddCmd(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		wantValue float64
		wantErr   bool
	}{
		{"No Argument", []string{}, 0, true},
	}

    for _, tt := range tests
}
