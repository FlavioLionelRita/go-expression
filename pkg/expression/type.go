package expression

import (
	"reflect"
)

type Kind byte

func (this Kind) Max(value Kind) Kind {
	if this > value {
		return this
	}
	return value
}

const (
	Invalid Kind = iota
	Void
	Int
	Float
	String
	Bool
	Object
	Array
	Any
)

type Value struct {
	value interface{}
	kind  Kind
}

func (this *Value) Type() Kind {
	return this.kind
}
func (this *Value) ToInt() int {
	return this.value.(int)
}
func (this *Value) ToFloat() float64 {
	return this.value.(float64)
}
func (this *Value) ToString() string {
	return this.value.(string)
}
func (this *Value) ToBool() bool {
	return this.value.(bool)
}
func (this *Value) ToObject() interface{} {
	return this.value
}

// https://stackoverflow.com/questions/44319906/why-golang-struct-array-cannot-be-assigned-to-an-interface-array
// https://research.swtch.com/interfaces
// https://gist.github.com/pmn/5374494
func (this *Value) toArray() []interface{} {
	list := reflect.ValueOf(this.value)
	var items []interface{}
	for i := 0; i < list.Len(); i++ {
		items = append(items, list.Index(i).Interface())
	}
	return items
}

func (this *Value) Len() int {
	//TODO:
	return 0
}
func (this *Value) IsEmpty() bool {

	switch this.kind {
	case Int:
		return this.ToInt() == 0
	case Float:
		return this.ToFloat() == 0
	case String:
		return this.ToString() == ""
	case Bool:
		return this.ToBool()
	case Object:
		return this.Len() == 0
	case Array:
		return this.Len() == 0
	}
	// https://glucn.medium.com/golang-an-interface-holding-a-nil-value-is-not-nil-bb151f472cc7
	_value := reflect.ValueOf(this.value)
	if _value.Kind() == reflect.Chan || _value.Kind() == reflect.Func || _value.Kind() == reflect.Map || _value.Kind() == reflect.Ptr || _value.Kind() == reflect.Slice {
		return _value.IsNil()
	}
	return _value.IsZero()
}

func ReflecValueToValue(source reflect.Value) *Value {
	//TODO
	return nil
}
