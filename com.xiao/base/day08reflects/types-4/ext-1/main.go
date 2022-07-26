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

	s1 := dataToStruct(PeopleData[0])
	fmt.Printf("%#v\n", s1)
}
/**
    数据库信息转换成结构体
 */
func dataToStruct(data []TableField) interface{} {
	var structFields []reflect.StructField
	structFields =make([]reflect.StructField,len(data))
	for i, tf := range data {
		structFields[i].Name=tf.Name
	}
	  typ:=reflect.StructOf(structFields)
      v:=reflect.New(typ).Elem()
	  s:=v.Addr().Interface()
	  return s
}