package main

import (
	repl "funlang/internal/prompt"
	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}
