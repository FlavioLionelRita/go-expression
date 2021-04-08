package helper

import (
	"reflect"
)

func In(x string, a []string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
func Max(a byte, b byte) byte {
	if a >= b {
		return a
	}
	return b
}

func Clone(source interface{}) interface{} {
	nInter := reflect.New(reflect.TypeOf(source).Elem())
	val := reflect.ValueOf(source).Elem()
	nVal := nInter.Elem()
	for i := 0; i < val.NumField(); i++ {
		nvField := nVal.Field(i)
		nvField.Set(val.Field(i))
	}
	return nInter.Interface()
}
