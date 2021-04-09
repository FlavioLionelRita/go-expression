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
	reAlphanumeric  *regexp.Regexp
	reInt           *regexp.Regexp
	reFloat         *regexp.Regexp
}

var singleton *ParserManager

func init() {
	singleton = &ParserManager{}
	singleton.reAlphanumeric, _ = regexp.Compile("p([a-zA-Z0-9_.]+)ch")
	singleton.reInt, _ = regexp.Compile("p([0-9]+)ch")
	singleton.reFloat, _ = regexp.Compile("p([0-9]+(.[0-9]*)?|.[0-9]+)([eE][0-9]+)ch")
	singleton.initOperator()
	singleton.initFunctions()
	singleton.initEnums()
	singleton.Refresh()
}
func GetParser() *ParserManager {
	return singleton
}

func (this *ParserManager) initOperator() {
	addition := Addition{Operator{name: "addition", category: "arithmetic", priority: 4}}
	subtraction := Subtraction{Operator{name: "subtraction", category: "arithmetic", priority: 4}}

	this.AddOperator("+", &addition)
	this.AddOperator("-", &subtraction)
}
func (this *ParserManager) initFunctions() {
}
func (this *ParserManager) initEnums() {
}
func (this *ParserManager) Refresh() {
}

func (this *ParserManager) AddOperator(key string, value interface{}) {
	this.operators[key] = value
}

func (this *ParserManager) newOperator(key string, oper1 *IOperand, oper2 *IOperand) interface{} {
	template := this.operators[key].(IOperator)
	clone := helper.Clone(template).(IOperator)
	clone.SetOper1(*oper1)
	clone.SetOper2(*oper2)
	return &clone
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
func (this *ParserManager) Eval(operand *IOperand, context *Context) (interface{}, error) {
	this.setContext(operand, context)
	result, err := (*operand).Value()
	if err != nil {
		return nil, err
	}
	return result.value, nil
}
func (this *ParserManager) Solve(source string, context *Context) (interface{}, error) {
	operand, err := this.Parse(source)
	if err != nil {
		return nil, err
	}
	return this.Eval(operand, context)
}
