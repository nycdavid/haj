package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"haj/lexer"
	"haj/parser"
	"haj/token"
)

const PROMPT = ">>> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if line == "exit" {
			os.Exit(0)
		}

		l := lexer.New(line)
		p := parser.New(l)

		p.ParseProgram()

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
