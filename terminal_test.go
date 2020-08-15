package go_terminal_module

import (
	"fmt"
	"testing"
)

func TestExecCommand(t *testing.T) {
	err, stderr, stdout := ExecCommand("echo 'Hello world'")
	PrintStderr(stderr)

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(stdout)
}

func TestExecCommandFulfillment(t *testing.T) {
	ExecCommandFulfillment("echo 'Hello world'")
}
