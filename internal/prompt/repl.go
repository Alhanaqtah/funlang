package repl

import (
	"bufio"
	"fmt"
	"io"

	"lang/internal/lexer"
	"lang/internal/token"
)

const prompt = "\033[1;36m➜ \033[0m"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		print(prompt)

		if ok := scanner.Scan(); !ok {
			return
		}

		input := scanner.Text()

		l := lexer.New(input)

		for tok := l.Next(); tok.Type != token.EOF; tok = l.Next() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
