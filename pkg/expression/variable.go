package expression

type Variable struct {
	name    string
	context Context
}

func (p *Variable) Value() interface{} {
	return p.context.Value(p.name)
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
