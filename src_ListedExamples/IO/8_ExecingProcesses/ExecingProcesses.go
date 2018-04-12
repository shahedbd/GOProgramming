package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	//cmd := exec.Command("sleep", "1")
	cmd := exec.Command("go", "build")
	//cmd := exec.Command("go", "version")

	log.Printf("Running command and waiting for it to finish...")
	err := cmd.Run()
	log.Printf("Command finished with error: %v", err)
	//printError(err)
}

func printError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
	}
}
