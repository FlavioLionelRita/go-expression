package expression

import (
	"errors"
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
	kind := oper1.kind.Max(oper2.kind)
	switch kind {
	case Int:
		return &Value{value: oper1.ToInt() - oper2.ToInt(), kind: kind}, nil
	case Float:
		return &Value{value: oper1.ToFloat() - oper2.ToFloat(), kind: kind}, nil
	case String:
		return &Value{value: oper1.ToString() + oper2.ToString(), kind: kind}, nil
	case Bool:
		return &Value{value: oper1.ToBool() && oper2.ToBool(), kind: kind}, nil
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

	kind := oper1.kind.Max(oper2.kind)
	switch kind {
	case Int:
		return &Value{value: oper1.ToInt() - oper2.ToInt(), kind: kind}, nil
	case Float:
		return &Value{value: oper1.ToFloat() - oper2.ToFloat(), kind: kind}, nil
	default:
		return nil, errors.New("invalid type for operator " + this.name)
	}
}
