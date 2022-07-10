package main

import (
	"fmt"
	"reflect"
)

/**
 * Created by safoti on 2022/7/9 21:42
   反射获取与设置值
**/
type Student struct {
	Id int
	Name string
	Weight float32
	Height float32
}
func main() {

   iValue:=reflect.ValueOf(1)
	ssValue:=reflect.TypeOf("string")
	userPtrValue :=reflect.ValueOf(Student{Id: 2,
		Name: "zs",
		Weight: 65,
		Height: 1.68,
	})

	fmt.Println(iValue)       //1
	fmt.Println(ssValue)       //hello
	fmt.Println(userPtrValue)


	//value --->type
    iType:=iValue.Type()
	userType := userPtrValue.Type()
	fmt.Println(iType.Kind() == reflect.Int, iValue.Kind() == reflect.Int, iType.Kind() == iValue.Kind())
	fmt.Println(userType.Kind() == reflect.Ptr, userPtrValue.Kind() == reflect.Ptr, userType.Kind() == userPtrValue.Kind())

//https://zhuanlan.zhihu.com/p/411313885
  userValues:= userPtrValue.Elem()
  fmt.Println(userValues.Kind(),userPtrValue.Kind())
}

