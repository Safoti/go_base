package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/26
 * @Description
 **/
func main() {
	dumpData(PeopleData)

	s1 := dataToStructs(PeopleData[0])
	fmt.Printf("%#v\n", s1)
}
func dataToStructs(data []TableField) interface{} {
	var structFields []reflect.StructField
	// number of struct fields required
	structFields = make([]reflect.StructField, len(data))

	var t reflect.Type

	for i, tf := range data {
		structFields[i].Name = tf.Name

		switch tf.Type {
		case "int":
			t = reflect.TypeOf(int(0))
		case "uint8":
			t = reflect.TypeOf(uint(0))
		case "string":
			t = reflect.TypeOf(string(""))
		case "float64":
			t = reflect.TypeOf(float64(0.0))
		}
		structFields[i].Type = t
	}

	typ := reflect.StructOf(structFields)
	v := reflect.New(typ).Elem()
	s := v.Addr().Interface()
	return s
}