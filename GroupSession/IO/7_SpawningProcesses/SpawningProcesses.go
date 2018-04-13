package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	//cmd := exec.Command("ls", "1")
	cmd := exec.Command("go", "version")

	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput

	printCommand(cmd)
	err := cmd.Run()
	printError(err)
	printOutput(cmdOutput.Bytes())
}

func printCommand(cmd *exec.Cmd) {
	fmt.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}

func printOutput(outs []byte) {
	if len(outs) > 0 {
		fmt.Printf("==> Output: %s\n", string(outs))
	}
}
