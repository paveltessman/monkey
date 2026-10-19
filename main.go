package main

import (
	"fmt"
	"os"

	"github.com/paveltessman/monkey/repl"
)


func main() {
	fmt.Println("Monkey programming language!")
	repl.Start(os.Stdin, os.Stdout)
}
