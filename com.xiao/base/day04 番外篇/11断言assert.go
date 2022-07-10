package main

import "fmt"

type National struct {
	Name string
	Code  int64
	Address string
}

type  Earth struct {
	National
}

func(n *National) sayNational(name string)  {
	fmt.Println("my country is : ",name)
}
func main() {
 //n:=National{Name: "US",Code: 10000,Address: "NORTH AMERICA"}
 n:=Earth{National{Name: "US",Code: 10000,Address: "NORTH AMERICA"}}
 check(n)
}
//类型判断（也叫断言）
func  check(v interface{})  {
	//fmt.Println(v.(National).Name)
	switch v.(type) {
	case National:
		fmt.Println(v.(National).Name)
	case Earth:
		fmt.Println(v.(Earth).Name)
	}
}