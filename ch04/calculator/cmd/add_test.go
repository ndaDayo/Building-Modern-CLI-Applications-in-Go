package cmd

import (
	"bytes"
	"os"
	"testing"
)

func setArgs(args []string) {
	os.Args = append([]string{os.Args[0]}, args...)
}

func TestAddCmdRun(t *testing.T) {
	cmd := addCmd

	var stdout bytes.Buffer
	cmd.SetOut(&stdout)

	cmd.SetArgs([]string{})

	err := cmd.Execute()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	expectedOutput := "Hello, John!\n"
	if stdout.String() != expectedOutput {
		t.Errorf("Expected output: %q, but got: %q", expectedOutput, stdout.String())
	}
}
