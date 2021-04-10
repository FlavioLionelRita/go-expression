package expression

// func nvl(value interface{}, _default interface{}) (interface{}, error) {
// 	rvalue:= reflect.ValueOf(value)

// }

// type Int struct{}

func nvl(value *Value, _default *Value) (*Value, error) {

	if value.IsEmpty() {
		return _default, nil
	}
	return value, nil
}
