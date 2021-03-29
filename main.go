package main

import (
	"fmt"

	exp "main/expression"
)

func main() {
	parser := exp.NewParser()
	fmt.Println(string(parser.Parse("a + b")))
}
