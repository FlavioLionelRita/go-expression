package model

import (
	core "main/pkg/core"
)

type Variable struct {
	name    string
	names   []string
	context core.Context
}

// func (p *Variable) Value() interface{} {
// 	var _value = p.context
// 	length := len(p.names)
// 	for i := 0; i < length; i++ {
// 		if i == length -1 {
// 			return _value[p.names[i]]
// 		}else{
// 			_value = _value[p.names[i]]
// 		}

// 	}
// }

func (p *Variable) Name() string {
	return p.name
}

func (p *Variable) Context() core.Context {
	return p.context
}
func (p *Variable) SetContext(value core.Context) {
	p.context = value
}
