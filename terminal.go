package terminal

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

var maxArgumentsError = "Only command and shell values are valid as arguments."

const (
	// BashShell Unix shell
	BashShell = "/bin/bash"
	// ShShell command language interpreter that executes commands
	ShShell = "/bin/sh"
)

// ExecCommand Máximun params are 2
// First param is the command
// Second is optional, select the shell
// return console output
func ExecCommand(s ...string) (string, string, error) {
	terminal := BashShell
	switch length := len(s); {
	case length > 2:
		log.Panic(maxArgumentsError)
	case length == 2:
		terminal = s[1]
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command(terminal, "-c", s[0])
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}

// ExecCommandNoReturn Máximun params are 2
// First param is the command
// Second is optional, select the shell
func ExecCommandNoReturn(s ...string) {
	var stdout, stderr string
	var err error
	switch length := len(s); {
	case length > 2:
		log.Panic(maxArgumentsError)
	case length == 2:
		stdout, stderr, err = ExecCommand(s[0], s[1])
		break
	case length == 1:
		stdout, stderr, err = ExecCommand(s[0])
		break
	}

	PrintStderr(stderr)
	fmt.Println(stdout)

	if err != nil {
		log.Panic(err)
	}
}

// PrintStderr print command output error
func PrintStderr(stderr string) {
	strings.TrimSuffix(stderr, "\n")
	if len(stderr) > 0 {
		fmt.Println(stderr)
	}
}
