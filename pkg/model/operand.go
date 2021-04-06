package model

type Operand interface {
	Value() (result interface{}, err error)
}

type _Operand struct {
}
