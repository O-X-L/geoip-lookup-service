package u

import (
	"reflect"
)

/*
func GetAttribute(dataStructure interface{}, name string) reflect.Value {
	return reflect.Indirect(
		reflect.ValueOf(&dataStructure),
	).Elem().FieldByName(name)
}
*/

func GetMapValue(dataStructure interface{}, name string) interface{} {
	return reflect.Indirect(
		reflect.ValueOf(&dataStructure),
	).Elem().Interface().(map[string]interface{})[name]
}
