package expression

var Type = newTypes()

func newTypes() *Types {
	return &Types{
		Int:    1,
		Float:  2,
		String: 3,
		Bool:   4,
		Object: 5,
		Array:  6,
		Any:    7,
	}
}

type Types struct {
	Int    byte
	Float  byte
	String byte
	Bool   byte
	Object byte
	Array  byte
	Any    byte
}

type Value struct {
	value interface{}
	_type byte
}

func (this *Value) Type() byte {
	return this._type
}

func (this *Value) toInt() int {
	return this.value.(int)
}
func (this *Value) toFloat() float64 {
	return this.value.(float64)
}
func (this *Value) toString() string {
	return this.value.(string)
}
func (this *Value) toBool() bool {
	return this.value.(bool)
}

type IOperand interface {
	Value() (*Value, error)
}

type Operand struct {
}

type KeyValue struct {
	name  string
	value IOperand
}

func (this *KeyValue) Name() string {
	return this.name
}
func (this *KeyValue) Value() (*Value, error) {
	return this.value.Value()
}

type Object struct {
	attributes []*KeyValue
}

func (this *Object) Value() (*Value, error) {
	dic := make(map[string]*Value)
	for i := 0; i < len(this.attributes); i++ {
		attribute := this.attributes[i]
		value, _error := attribute.Value()
		if _error != nil {
			return nil, _error
		}
		dic[attribute.name] = value
	}
	return &Value{value: dic, _type: Type.Object}, nil

}
