package main

import (
	"fmt"
	"log"
	"os/exec"
)

func execute(args string) {
	cmd := exec.Command("sh", "-c", args)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	if len(stdoutStderr) > 0 {
		fmt.Printf("%s\n", stdoutStderr)
	}
}
