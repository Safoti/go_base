package main

import (
	"fmt"
	"reflect"
)

/**
     反射部分后续在仔细阅读
 */


type User struct {
	Id int
	Name string
	Age int
}

func (u User)Hello()  {
	fmt.Println("Hello reflect")
}
func main() {
	u:=User{1,"lk",20}
	info(u)
}
/**
   反射处理
 */
func  info(o interface{})  {
  t:= reflect.TypeOf(o)
  fmt.Println(t.Name())
  v:= reflect.ValueOf(o)

	for i := 0; i <t.NumField() ; i++ {
		f:=t.Field(i)
		valls:=v.Field(i).Interface()
		fmt.Printf("%s: %v:=%v\n",f.Name,f.Type,valls)
	}
}