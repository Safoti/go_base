package main

import (
	"fmt"
)
/**
   番外编:
      通过源码来加深对golang 的语法的学习
 */
/**

函数与方法的区别：

  函数
     格式：func 函数名(函数入参) 函数出参{}

func Trim(s, cutset string) string {
	if s == "" || cutset == "" {
		return s
	}
	if len(cutset) == 1 && cutset[0] < utf8.RuneSelf {
		return trimLeftByte(trimRightByte(s, cutset[0]), cutset[0])
	}
	if as, ok := makeASCIISet(cutset); ok {
		return trimLeftASCII(trimRightASCII(s, &as), &as)
	}
	return trimLeftUnicode(trimRightUnicode(s, cutset), cutset)
}

  //方法
  func 传递的对象   函数名 （入参）出参 {}
func (as *asciiSet) contains(c byte) bool {
	return (as[c/32] & (1 << (c % 32))) != 0
}
 */
func main() {
	//s:="hello.name"
   //strings.Split(s,".")


	//测试下方法的调用

	lz:=lktbz{"lk",15}
	name := lz.getName(15)
	fmt.Println(name)
}

func MyFuncDemo(fi,se string)string{
	if fi=="" ||se ==""{
		return fi
	}
	return ""
}
//传递一个函数进来当参数
func TrimLeftFunc(s string, f func(rune) bool) string {

	return ""
}


//定义一个方法
func (u *lktbz)getName(age int)string{
	//check u
    //chekc age

	if age ==u.Age{
		fmt.Println(u.Name)
	   return u.Name
	}
	return ""
}
type lktbz struct {
	 Name string
	 Age int
}