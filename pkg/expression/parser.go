package expression

import (
	"errors"
	"main/pkg/helper"
	"reflect"
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

func (this *Parser) Parse() (interface{}, error) {
	var operands []*IOperand
	for !this.end() {
		operand, _error := this.getExpression(nil, "", ";")
		if _error != nil {
			return nil, _error
		}
		if operand == nil {
			break
		}
		operands = append(operands, operand.(*IOperand))
	}
	if len(operands) == 1 {
		return operands[0], nil
	}
	return &Block{lines: operands}, nil
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

func (this *Parser) getExpression(operand1 *IOperand, operator string, _break string) (interface{}, error) {
	var expression *IOperand
	var operand2 *IOperand
	var _error error
	isbreak := false
	for !this.end() {
		if operand1 == nil && operator == "" {
			operand1, _error = this.getOperand()
			if _error != nil {
				return nil, _error
			}
			operator = this.getOperator()
			if operator == "" || strings.Contains(_break, operator) {
				expression = operand1
				isbreak = true
				break
			}
			operand2, _error = this.getOperand()
			if _error != nil {
				return nil, _error
			}
			nextOperator := this.getOperator()
			if nextOperator == "" || strings.Contains(_break, nextOperator) {
				expression = this.manager.newOperator(operator, operand1, operand2).(*IOperand)
				isbreak = true
				break
			} else if this.priority(operator) >= this.priority(nextOperator) {
				operand1 = this.manager.newOperator(operator, operand1, operand2).(*IOperand)
				operator = nextOperator
			} else {
				newOperand, _error := this.getExpression(operand2, nextOperator, _break)
				if _error != nil {
					return nil, _error
				}
				expression = this.manager.newOperator(operator, operand1, newOperand.(*IOperand)).(*IOperand)
				isbreak = true
				break
			}
		}
		if !isbreak {
			expression = this.manager.newOperator(operator, operand1, operand2).(*IOperand)
		}
		// if all the operands are constant, reduce the expression a constant
		if expression != nil {
			p, hasOperands := (*expression).(IComposite)
			if hasOperands {
				allConstants := true
				for i := 0; i < len(p.operands()); i++ {
					operand := p.getOperand(i)
					if reflect.TypeOf(operand).Name() != "Constant" {
						allConstants = false
						break
					}
				}
				if allConstants {
					value, _error := (*expression).Value()
					if _error != nil {
						return nil, _error
					}
					return &Constant{value: value.value, _type: value._type}, nil
				}
			}
		}
		return expression, nil
	}
	return nil, nil
}
func (this *Parser) getOperand() (*IOperand, error) {
	//TODO
	return nil, nil
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
func (this *Parser) getArgs(end byte) ([]*IOperand, error) {
	endExpression := string([]byte{',', end})
	var args []*IOperand
	for !this.end() {
		arg, _error := this.getExpression(nil, "", endExpression)
		if _error != nil {
			return nil, _error
		}
		if arg != nil {
			args = append(args, arg.(*IOperand))
		}
		if this.previous() == end {
			break
		}
	}
	return args, nil
}
func (this *Parser) getObject(end byte) (*Object, error) {
	var attributes []*KeyValue
	for !this.end() {
		name := this.getValue()
		if this.current() == ':' {
			this.index++
		} else {
			return nil, errors.New("attribute " + name + " without value")
		}
		value, _error := this.getExpression(nil, "", ",}")
		if _error != nil {
			return nil, _error
		}
		attribute := KeyValue{name: name, value: value.(*IOperand)}
		attributes = append(attributes, &attribute)
		if this.previous() == end {
			this.index++
			break
		}
	}
	return &Object{attributes: attributes}, nil
}
