package cmd

import (
	"bytes"
	"calculator/storage"
	"strings"
	"testing"
)

func TestAddCmdRun(t *testing.T) {
	initialValue := 0.0
	storage.SetValue(initialValue)

	tests := []struct {
		name    string
		args    []string
		wantOut string
	}{
		{
			name:    "No Arguments",
			args:    []string{},
			wantOut: "command requires input value\n",
		},
		{
			name:    "Multiple Arguments",
			args:    []string{"10", "20"},
			wantOut: "only accepts a single argument\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := addCmd
			b := bytes.NewBufferString("")
			cmd.SetOut(b)
			cmd.SetErr(b)
			cmd.SetArgs(tt.args)

			cmd.Execute()

			if !strings.Contains(b.String(), tt.wantOut) {
				t.Errorf("TestAddCmd %s: got output = %v, want %v", tt.name, b.String(), tt.wantOut)
			}
		})
	}
}
