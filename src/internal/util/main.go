package u

import (
	"log"
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

func LogError(prefix string, err interface{}) {
	log.Printf("ERROR: %v | %v\n", prefix, err)
}

func LogWarn(prefix string, err interface{}) {
	log.Printf("WARNING: %v | %v\n", prefix, err)
}
