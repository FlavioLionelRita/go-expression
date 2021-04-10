package expression

type Constant struct {
	value *Value
}

func (this *Constant) Value() (*Value, error) {
	return this.value, nil
}
