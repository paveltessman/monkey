package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/paveltessman/monkey/lexer"
	"github.com/paveltessman/monkey/token"
)

const prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {

		fmt.Printf(prompt)

		scanned := scanner.Scan()
		err := scanner.Err()
		if err != nil {
			panic(err)
		}

		if !scanned {
			return
		}

		line := scanner.Text()

		lxr := lexer.New(line)

		for tok := lxr.NextToken(); tok.Type != token.EOF; tok = lxr.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
