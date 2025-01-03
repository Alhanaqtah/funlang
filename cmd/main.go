package main

import (
	repl "lang/internal/prompt"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
