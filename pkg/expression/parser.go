package expression

import (
	"errors"
	"main/pkg/helper"
	"strings"
)

type Parser struct {
	manager *ParserManager
	chars   []byte
	index   int
	length  int
}

func NewParser(manager *ParserManager, source string) *Parser {
	p := Parser{manager: manager}
	p.chars = p.clear(source)
	p.index = 0
	p.length = len(p.chars)

	return &p
}

func (this *Parser) Parse() (IOperand, error) {
	//TODO
	// return string(p.chars)
	return nil, nil
}

func (this *Parser) clear(source string) []byte {
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
func (this *Parser) previous() byte {
	return this.chars[this.index-1]
}
func (this *Parser) current() byte {
	return this.chars[this.index]
}
func (this *Parser) next() byte {
	return this.chars[this.index+1]
}
func (this *Parser) end() bool {
	return this.index >= this.length
}

func (this *Parser) getExpression(operand1 *IOperand, operator *IOperator, _break string) *IOperand {
	//TODO
	return nil
}
func (this *Parser) getOperand() IOperand {
	//TODO
	return nil
}
func (this *Parser) priority(op string) byte {
	return this.manager.priority(op)
}

func (this *Parser) getValue() string {
	buff := make([]byte, 50)
	j := 0
	for ; !this.end() && this.manager.reAlphanumeric.Match([]byte{this.current()}); j++ {
		if j >= 50 {
			buff = append(buff, this.current())
		} else {
			buff[j] = this.current()
		}
		this.index++
	}
	return string(buff[:j])
}
func (this *Parser) getOperator() string {
	if this.end() {
		return ""
	}
	op := ""
	if this.index+2 < this.length {
		triple := string([]byte{this.current(), this.next(), this.chars[this.index+2]})
		if helper.In(triple, this.manager.tripleOperators) {
			op = triple
		}
	}
	if op == "" && this.index+1 < this.length {
		double := string([]byte{this.current(), this.next()})
		if helper.In(double, this.manager.doubleOperators) {
			op = double
		}
	}
	if op == "" {
		op = string(this.current())
	}
	this.index += len(op)
	return op
}
func (this *Parser) getString(char byte) string {
	var sb strings.Builder
	for !this.end() {
		if this.current() == char {
			if !((this.index+1 < this.length && this.next() == char) || (this.previous() == char)) {
				break
			}
		}
		sb.WriteByte(this.current())
		this.index++
	}
	this.index++
	return sb.String()
}
func (this *Parser) getArgs(end byte) []*IOperand {
	endExpression := string([]byte{',', end})
	var args []*IOperand
	for !this.end() {
		arg := this.getExpression(nil, nil, endExpression)
		if arg != nil {
			args = append(args, arg)
		}
		if this.previous() == end {
			break
		}
	}
	return args
}
func (this *Parser) getObject(end byte) (Object, error) {
	var attributes []*KeyValue
	for !this.end() {
		name := this.getValue()
		if this.current() == ':' {
			this.index++
		} else {
			return Object{}, errors.New("attribute " + name + " without value")
		}
		value := this.getExpression(nil, nil, ",}")
		attribute := KeyValue{name: name, value: *value}
		attributes = append(attributes, &attribute)
		if this.previous() == end {
			this.index++
			break
		}
	}
	return Object{attributes: attributes}, nil
}
