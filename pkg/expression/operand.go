package expression

var Type = newTypes()

func newTypes() *Types {
	return &Types{
		Int:     1,
		Float:   2,
		String:  3,
		Boolean: 4,
		Object:  5,
		Array:   6,
		Any:     7,
	}
}

type Types struct {
	Int     byte
	Float   byte
	String  byte
	Boolean byte
	Object  byte
	Array   byte
	Any     byte
}

type Value struct {
	value interface{}
	_type byte
}

func (this *Value) Type() byte {
	return this._type
}

type IOperand interface {
	Value() (*Value, error)
}

type Operand struct {
}
