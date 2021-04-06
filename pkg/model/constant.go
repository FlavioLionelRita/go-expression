package model

type Constant struct {
	value string
	_type string
	_Operand
}

func (p *Constant) Value() interface{} {
	//TODO: ver como convertir de acuerdo al tipo
	return p.value
}
func (p *Constant) Type() string {
	return p._type
}
