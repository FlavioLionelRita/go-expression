package core

type Context struct {
	values map[string]interface{}
}

func (p *Context) Value(name string) interface{} {
	//TODO: en name , puede venir  a.b.c  hay que resolver como acceder a propiedades de acuerdo al path
	return p.values[name]

	// length := len(p.names)
	// for i := 0; i < length; i++ {
	// 	if i == length-1 {
	// 		return _value[p.names[i]]
	// 	} else {
	// 		_value = _value[p.names[i]]
	// 	}
	// }
}

func (p *Context) SetValue(name string, value interface{}) {
	//TODO: en name , puede venir  a.b.c  hay que resolver como acceder a propiedades de acuerdo al path
	p.values[name] = value
}
