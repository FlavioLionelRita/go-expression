package expression

import (
	"main/lib"
	"regexp"
)

type Parser struct {
	chars          []byte
	index          int
	length         int
	reAlphanumeric *regexp.Regexp
}

func NewParser() Parser {
	return Parser{}
}

func (p *Parser) Parse(source string) string {
	p.chars = p.clear(source)
	p.index = 0
	p.length = len(p.chars)
	p.reAlphanumeric, _ = regexp.Compile("p([a-zA-Z0-9_.]+)ch")

	//TODO
	return string(p.chars)
}

func (p *Parser) clear(source string) []byte {
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
	return result[:j]
}
func (p *Parser) previous() byte {
	return p.chars[p.index-1]
}
func (p *Parser) current() byte {
	return p.chars[p.index]
}
func (p *Parser) next() byte {
	return p.chars[p.index+1]
}
func (p *Parser) end() bool {
	return p.index >= p.length
}
func (p *Parser) priority(op string) int {
	if lib.In(op, []string{"=", "+=", "-=", "*=", "/=", "%=", "**=", "//=", "&=", "|=", "^=", "<<=", ">>="}) {
		return 1
	}
	if lib.In(op, []string{"&&", "||"}) {
		return 2
	}
	if lib.In(op, []string{">", "<", ">=", "<=", "!=", "=="}) {
		return 3
	}
	if lib.In(op, []string{"+", "-"}) {
		return 4
	}
	if lib.In(op, []string{"*", "/"}) {
		return 5
	}
	if lib.In(op, []string{"**", "//"}) {
		return 6
	}
	return -1
}
func (p *Parser) getValua() string {
	buff := make([]byte, 50)
	j := 0
	for ; p.end() == false; j++ {
		buff[j] = p.current()
		p.index++
	}
	return string(buff[:j])

}
