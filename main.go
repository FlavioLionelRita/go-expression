package main

import (
	"fmt"

	exp "main/pkg/expression"
)

func main() {

	result := exp.Parse("a + b")
	xType := fmt.Sprintf("%v", result)
	fmt.Println(string(xType))
}
