package main

import (
	"os"

	"github.com/ryangerardwilson/musica/internal/cli"
)

func main() {
	os.Exit(cli.Command(os.Args[1:], os.Stdin, os.Stdout, os.Stderr))
}
