package main

import (
	"fmt"
	"os"

	"haj/repl"
)

func main() {
	fmt.Println("Haj 0.1")
	repl.Start(os.Stdin, os.Stdout)
}
