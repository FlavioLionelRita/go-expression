// You can edit this code!
// Click here and start typing.
package main

import "fmt"

type Parser struct {
	chars  string
	index  int
	length int
}

func (p *Parser) clear(source string) string {
	length := len(source)
	j := 0
	var quotes byte
	var isString bool = false
	var result = make([]byte, length)
	for i := 0; i < length; i++ {
		p := source[i]
		if isString && p == quotes {
			isString = false
		} else if !isString && (p == '\'' || p == '"') {
			isString = true
			quotes = p
		}
		if p != ' ' || isString {
			result[j] = p
			j++
		}
	}
	return string(result[:j])

}

func main() {
	parser := Parser{}
	var a string
	a = parser.clear("a + b")
	fmt.Println(a)
}
