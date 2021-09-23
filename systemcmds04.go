package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"strings"

)

func upper(data string) string {
	return strings.ToUpper(data)
}

func main() {
	cmd := exec.Command("echo", "the third moon of endor")

	// our command will output through this pipe
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		log.Fatal(err)
	}

	// if we use start, we must use wait to ensure our command will finish
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	// grab the stdout from our command and read it in as string data
	data, err := ioutil.ReadAll(stdout)

	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}

	// display data
	fmt.Printf("%s\n", upper(string(data)))
}
