package go_terminal_module

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func ExecCommand(command string) (error, string, string) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("/bin/bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	return err, stdout.String(), stderr.String()
}

func ExecCommandFulfillment(command string) {
	err, stdout, stderr := ExecCommand(command)
	PrintStderr(stderr)
	fmt.Println(stdout)

	if err != nil {
		log.Panic(err)
	}
}

func PrintStderr(stderr string) {
	strings.TrimSuffix(stderr, "\n")
	if len(stderr) > 0 {
		fmt.Println(stderr)
	}
}

