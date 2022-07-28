package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description   反射标准库 之structof 学习
 **/
func main() {
  typ:=reflect.StructOf([]reflect.StructField{
  		{
  			Name: "Height",   //构建字段名与类型
  			Type: reflect.TypeOf(float64(0)),
  			Tag: `json:Height`,
		},{
		  Name: "Age",
		  Type: reflect.TypeOf(int(0)),
		  Tag:  `json:"age"`,
	  },
	})

	v :=  reflect.New(typ).Elem()
	v.Field(0).SetFloat(0.1)
	v.Field(1).SetInt(2)
	s:=v.Addr().Interface()

	w := new(bytes.Buffer)
	if err := json.NewEncoder(w).Encode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)
	fmt.Printf("json:  %s", w.Bytes())


	r := bytes.NewReader([]byte(`{"height":1.5,"age":10}`))
	if err := json.NewDecoder(r).Decode(s); err != nil {
		panic(err)
	}
	fmt.Printf("value: %+v\n", s)
}
