package cmd

import (
	"testing"
)

func TestAddCmdRun(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantOut string
		wantErr bool
	}{
		{
			name:    "No Arguments",
			args:    []string{},
			wantOut: "command requires input value\n",
			wantErr: true,
		},
		{
			name:    "Multiple Arguments",
			args:    []string{"10", "20"},
			wantOut: "only accepts a single argument\n",
			wantErr: true,
		},
		{
			name:    "Invalid Argument",
			args:    []string{"abc"},
			wantOut: "unable to parse input[abc]",
			wantErr: true,
		},
		{
			name:    "Valid Argument",
			args:    []string{"10"},
			wantOut: "",
			wantErr: false,
		},
	}
}
