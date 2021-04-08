package expression

import (
	"errors"
	helper "main/pkg/helper"
)

type Addition struct {
	Operator
}

// func (this *Addition) clone(oper1 *IOperand, oper2 *IOperand) IOperator {
// 	clone := new(Addition)
// 	clone.name = this.name
// 	clone.category = this.category
// 	clone.oper1 = oper1
// 	clone.oper2 = oper2
// 	return clone

// }
func (this *Addition) Value() (*Value, error) {
	oper1, err := this.Oper1().Value()
	if err != nil {
		return nil, err
	}
	oper2, err := this.Oper2().Value()
	if err != nil {
		return nil, err
	}
	_type := helper.Max(oper1.Type(), oper2.Type())
	switch _type {
	case Type.Int:
		return &Value{value: oper1.toInt() - oper2.toInt(), _type: _type}, nil
	case Type.Float:
		return &Value{value: oper1.toFloat() - oper2.toFloat(), _type: _type}, nil
	case Type.String:
		return &Value{value: oper1.toString() + oper2.toString(), _type: _type}, nil
	case Type.Bool:
		return &Value{value: oper1.toBool() && oper2.toBool(), _type: _type}, nil
	default:
		return nil, errors.New("invalid type for operator " + this.name)
	}
}

type Subtraction struct {
	Operator
}

func (this *Subtraction) Value() (*Value, error) {
	oper1, err := this.Oper1().Value()
	if err != nil {
		return nil, err
	}
	oper2, err := this.Oper2().Value()
	if err != nil {
		return nil, err
	}
	_type := helper.Max(oper1.Type(), oper2.Type())
	switch _type {
	case Type.Int:
		return &Value{value: oper1.toInt() - oper2.toInt(), _type: _type}, nil
	case Type.Float:
		return &Value{value: oper1.toFloat() - oper2.toFloat(), _type: _type}, nil
	default:
		return nil, errors.New("invalid type for operator " + this.name)
	}
}
