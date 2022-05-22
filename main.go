package main

import (
	"os"

	"github.com/BradMyrick/GoBlockChain/cli"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}
	cmd.Run()
}
