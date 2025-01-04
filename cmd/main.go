package main

import (
	"os"

	"funlang/internal/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
