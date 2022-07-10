package main

import (
	"fmt"
	"reflect"
)

/**
 * Created by safoti on 2022/7/9 21:28
**/

type Dog struct {

}
func (t Dog) T1(nums []int) {
	fmt.Println(nums)
	fmt.Println("t1")
}
func (t Dog) T2() {
	fmt.Println("t2")
}
func (t *Dog)t3(){
	fmt.Println("t3")

}
func main() {
   var tt Dog
   //获取对象的方法
   getType:=reflect.TypeOf(tt)
	for i := 0; i <getType.NumMethod() ; i++ {
		met :=getType.Method(i)
		fmt.Printf("%s, %s, %s, %d\n", met.Type, met.Type.Kind(), met.Name, met.Index)

	}


//获取函数
  addFunc:=reflect.TypeOf(Add)
	fmt.Printf("is function type %t\n", addFunc.Kind() == reflect.Func)
	argInNum :=addFunc.NumIn()
	argOutNum :=addFunc.NumOut()

	for i := 0; i < argInNum; i++ {
		argTyp := addFunc.In(i)
		fmt.Printf("第%d个输入参数的类型%s\n", i, argTyp)
	}
	for i := 0; i < argOutNum; i++ {
		argTyp := addFunc.Out(i)
		fmt.Printf("第%d个输出参数的类型%s\n", i, argTyp)
	}

}

func Add(a,b int) int  {
	return a+b
}
