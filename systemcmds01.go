package main

import (
	"log"
	"os/exec"
)

func main() {

	// prepare a "cmd" strut
	cmd := exec.Command("ls")

	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}
}


