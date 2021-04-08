package expression

type IOperator interface {
	Name() string
	Category() string
	Priority() byte
	Oper1() IOperand
	Oper2() IOperand
	clone(oper1 *IOperand, oper2 *IOperand) IOperator
	// SetName(value string)
	// SetCategory(value string)
	SetOper1(value IOperand)
	SetOper2(value IOperand)
}

type Operator struct {
	name     string
	category string
	priority byte
	oper1    *IOperand
	oper2    *IOperand
}

func (this *Operator) Name() string {
	return this.name
}

// func (this *Operator) SetName(value string) {
// 	this.name = value
// }
func (this *Operator) Category() string {
	return this.category
}

// func (this *Operator) SetCategory(value string) {
// 	this.category = value
// }
func (this *Operator) Priority() byte {
	return this.priority
}
func (this *Operator) Oper1() IOperand {
	return *this.oper1
}

func (this *Operator) SetOper1(value *IOperand) {
	this.oper1 = value
}
func (this *Operator) Oper2() IOperand {
	return *this.oper2
}

func (this *Operator) SetOper2(value *IOperand) {
	this.oper2 = value
}
