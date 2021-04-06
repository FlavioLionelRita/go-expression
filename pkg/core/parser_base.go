package core

import (
	"main/pkg/helper"
	"regexp"
)

type _Parser struct {
	chars          []byte
	index          int
	length         int
	reAlphanumeric *regexp.Regexp
}

func NewParser(source string) *_Parser {
	p := _Parser{}
	p.chars = p.clear(source)
	p.index = 0
	p.length = len(p.chars)
	p.reAlphanumeric, _ = regexp.Compile("p([a-zA-Z0-9_.]+)ch")
	return &p
}

func (p *_Parser) Parse() string {
	//TODO
	return string(p.chars)
}

func (p *_Parser) clear(source string) []byte {
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
func (p *_Parser) previous() byte {
	return p.chars[p.index-1]
}
func (p *_Parser) current() byte {
	return p.chars[p.index]
}
func (p *_Parser) next() byte {
	return p.chars[p.index+1]
}
func (p *_Parser) end() bool {
	return p.index >= p.length
}
func (p *_Parser) priority(op string) int {
	if helper.In(op, []string{"=", "+=", "-=", "*=", "/=", "%=", "**=", "//=", "&=", "|=", "^=", "<<=", ">>="}) {
		return 1
	}
	if helper.In(op, []string{"&&", "||"}) {
		return 2
	}
	if helper.In(op, []string{">", "<", ">=", "<=", "!=", "=="}) {
		return 3
	}
	if helper.In(op, []string{"+", "-"}) {
		return 4
	}
	if helper.In(op, []string{"*", "/"}) {
		return 5
	}
	if helper.In(op, []string{"**", "//"}) {
		return 6
	}
	return -1
}
func (p *_Parser) getValue() string {
	buff := make([]byte, 50)
	j := 0
	for ; !p.end() && p.reAlphanumeric.Match([]byte{p.current()}); j++ {
		if j >= 50 {
			buff = append(buff, p.current())
		} else {
			buff[j] = p.current()
		}
		p.index++
	}
	return string(buff[:j])
}
func (p *_Parser) getOperator() string {
	if p.end() {
		return ""
	}
	op := ""
	if p.index+2 < p.length {
		triple := string([]byte{p.current(), p.next(), p.chars[p.index+2]})
		if helper.In(triple, []string{"**=", "//=", "<<=", ">>="}) {
			op = triple
		}
	}
	if op == "" && p.index+1 < p.length {
		double := string([]byte{p.current(), p.next()})
		if helper.In(double, []string{"**", "//", ">=", "<=", "!=", "==", "+=", "-=", "*=", "/=", "%=", "&&", "||", "|=", "^=", "<<", ">>"}) {
			op = double
		}
	}
	if op == "" {
		op = string(p.current())
	}
	p.index += len(op)
	return op
}
