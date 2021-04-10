package expression

import (
	"reflect"
)

type IOperand interface {
	Value() (*Value, error)
}

type IComposite interface {
	operands() []*IOperand
	getOperand(index int) *IOperand
}

type Operand struct {
}

type KeyValue struct {
	name  string
	value *IOperand
}

func (this *KeyValue) Name() string {
	return this.name
}
func (this *KeyValue) Value() (*Value, error) {
	return (*this.value).Value()
}

type _Object struct {
	attributes []*KeyValue
}

func (this *_Object) Value() (*Value, error) {
	dic := make(map[string]*Value)
	for i := 0; i < len(this.attributes); i++ {
		attribute := this.attributes[i]
		value, _error := attribute.Value()
		if _error != nil {
			return nil, _error
		}
		dic[attribute.name] = value
	}
	return &Value{value: dic, kind: Object}, nil

}

type Block struct {
	lines []*IOperand
}

func (this *Block) Value() (*Value, error) {
	var result *Value
	var line IOperand
	var _error error
	for i := 0; i < len(this.lines); i++ {
		line = *this.lines[i]
		result, _error = line.Value()
		if _error != nil {
			return nil, _error
		}
	}
	return result, nil
}

type IndexDecorator struct {
	variable *IOperand
	index    *IOperand
}

func (this *IndexDecorator) Value() (*Value, error) {
	_variable, _error := (*this.variable).Value()
	if _error != nil {
		return nil, _error
	}
	_index, _error := (*this.index).Value()
	if _error != nil {
		return nil, _error
	}
	list := reflect.ValueOf(_variable.value)
	result := list.Index(_index.ToInt())
	return ReflecValueToValue(result), nil
}
