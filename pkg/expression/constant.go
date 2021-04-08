package expression

type Constant struct {
	value interface{}
	_type byte
	Operand
}

func (this *Constant) Value() (*Value, error) {
	return &Value{value: this.value, _type: this._type}, nil
}
