package expression

import (
	"reflect"
)

type Variable struct {
	name    string
	context Context
}

func (p *Variable) Value() (*Value, error) {
	_value := p.context.Value(p.name)
	rvalue := reflect.ValueOf(_value)
	return ReflecValueToValue(rvalue), nil

}
func (p *Variable) SetValue(value interface{}) {
	p.context.SetValue(p.name, value)
}
func (p *Variable) Name() string {
	return p.name
}
func (p *Variable) Context() Context {
	return p.context
}
func (p *Variable) SetContext(value Context) {
	p.context = value
}
