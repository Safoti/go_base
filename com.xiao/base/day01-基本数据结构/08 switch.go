package main

import (
	"fmt"
	"math"
)

func switchDemo(name string,number int){
	switch name {
	case "coding3min":
		fmt.Println("welcome" + name)
	default:
		fmt.Println("403 forbidden:" + name)
		return
	}
	switch  {
	case number>=90:
		fmt.Println("优秀")
	case number >= 80:
		fmt.Println("良好")
	case number >= 60:
		fmt.Println("凑合")
	default:
		fmt.Println("太搓了")
	}


}
//骚操作
//type-switch 用来判断某个interface变量中实际存储的变量类型
func  typeSwitchDemo(x interface{})int  {
	switch t:=x.(type) {
	case int:
		return t
	case float64:
		return int(math.Ceil(t))
	}
	return 0
}
func main() {
   
}
