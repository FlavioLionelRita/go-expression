package expression

import (
	"errors"
	"reflect"
)

type Function struct {
	name     string
	operands []*IOperand
	isChild  bool
}

// https://www.programmersought.com/article/45971546129/
func (this *Function) Value() (*Value, error) {

	var count = len(this.operands)
	args := make([]reflect.Value, count)

	for i := 0; i < count; i++ {
		operand := this.operands[i]
		value, _error := (*operand).Value()
		if _error != nil {
			return nil, _error
		}
		args[i] = reflect.ValueOf(value.value)
	}

	_funct := reflect.ValueOf(parserManager.getFunction(this.name))
	if count != _funct.Type().NumIn() {
		return nil, errors.New("The number of params is not adapted.")
	}

	results := _funct.Call(args)
	if len(results) == 0 {
		return &Value{value: nil, kind: Void}, nil
	} else if len(results) == 1 {
		result := results[0]
		return ReflecValueToValue(result), nil
	} else {
		values := make([]*Value, len(results))
		for i := 0; i < len(results); i++ {
			result := results[i]
			values[i] = ReflecValueToValue(result)
		}
		return &Value{value: values, kind: Array}, nil
	}
}
