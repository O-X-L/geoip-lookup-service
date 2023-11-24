package u

import (
	"reflect"
)

func GetAttribute(dataStructure interface{}, name string) reflect.Value {
	return reflect.Indirect(
		reflect.ValueOf(&dataStructure),
	).FieldByName(name)
}
