package expression

import (
	"main/pkg/helper"
	"regexp"
)

type ParserManager struct {
	operators       map[string]interface{}
	functions       map[string]interface{}
	enums           map[string]interface{}
	doubleOperators []string
	tripleOperators []string
	isAlphanumeric  func(s string) bool
	isInt           func(s string) bool
	isFloat         func(s string) bool
}

var parserManager *ParserManager

func init() {
	parserManager = &ParserManager{}
	parserManager.isAlphanumeric = regexp.MustCompile(`^[a-zA-Z0-9_.]+$`).MatchString
	parserManager.isInt = regexp.MustCompile(`[0-9]+$`).MatchString
	parserManager.isFloat = regexp.MustCompile(`([0-9]+([0-9.]*)?|[0-9.]+)([eE][0-9]+)?`).MatchString
	parserManager.initOperator()
	parserManager.initFunctions()
	parserManager.initEnums()
	parserManager.Refresh()
}
func GetParser() *ParserManager {
	return parserManager
}

func (this *ParserManager) initOperator() {
	addition := Addition{Operator{name: "addition", category: "arithmetic", priority: 4}}
	subtraction := Subtraction{Operator{name: "subtraction", category: "arithmetic", priority: 4}}

	this.AddOperator("+", &addition)
	this.AddOperator("-", &subtraction)
}
func (this *ParserManager) initFunctions() {
	this.AddFunction("nvl", nvl)
}
func (this *ParserManager) initEnums() {
}
func (this *ParserManager) Refresh() {
}

func (this *ParserManager) AddOperator(key string, value interface{}) {
	this.operators[key] = value
}
func (this *ParserManager) AddFunction(key string, value interface{}) {
	this.functions[key] = value
}
func (this *ParserManager) AddEnum(key string, value interface{}) {
	this.enums[key] = value
}

func (this *ParserManager) newOperator(key string, oper1 *IOperand, oper2 *IOperand) IOperator {
	template := this.operators[key]
	clone := helper.Clone(template).(IOperator)
	clone.SetOper1(*oper1)
	clone.SetOper2(*oper2)
	return clone
	// clone := reflect.New(reflect.ValueOf(template).Elem().Type()).Interface().(IOperator)
	// clone.SetName(template.Name())
	// clone.SetCategory(template.Category())
	// clone.SetOper1(*oper1)
	// clone.SetOper2(*oper2)
	// return clone
}
func (this *ParserManager) priority(key string) byte {
	return this.operators[key].(IOperator).Priority()
}

func (this *ParserManager) getFunction(key string) interface{} {
	return this.functions[key]
}

func (this *ParserManager) setContext(operand *IOperand, context *Context) {
	//TODO
}

func (this *ParserManager) Parse(source string) (*IOperand, error) {
	parser := NewParser(this, source)
	operand, err := parser.Parse()
	if err != nil {
		return nil, err
	}
	return operand.(*IOperand), nil
}
func (this *ParserManager) Eval(operand *IOperand, context *Context) (*Value, error) {
	this.setContext(operand, context)
	result, err := (*operand).Value()
	if err != nil {
		return nil, err
	}
	return result, nil
}
func (this *ParserManager) Solve(source string, context *Context) (*Value, error) {
	operand, err := this.Parse(source)
	if err != nil {
		return nil, err
	}
	return this.Eval(operand, context)
}
