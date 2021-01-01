package terminal

import (
	"fmt"
	"testing"
)

func TestExecCommand(t *testing.T) {
	stderr, stdout, err := ExecCommand("echo 'Hello world'")
	PrintStderr(stderr)

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(stdout)
}

func TestExecCommandSh(t *testing.T) {
	stderr, stdout, err := ExecCommand("echo 'Hello world'", ShShell)
	PrintStderr(stderr)

	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println(stdout)
}

func TestExecCommandNoReturn(t *testing.T) {
	ExecCommandNoReturn("echo 'Hello world'")
}

func TestExecCommandNoReturnSh(t *testing.T) {
	ExecCommandNoReturn("echo 'Hello world'", ShShell)
}
