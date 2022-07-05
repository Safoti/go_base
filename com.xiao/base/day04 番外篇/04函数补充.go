package main

import "fmt"

/**
 函数值传递与引用传递demo

*/
//值传递例子
func  noChange(a,b int)  {
	tmp:=a
	a=b
	b=tmp
}
//引用传递例子
func noChangePoint(a,b  *int){
 tmp:=*a
 *a=*b
 *b=tmp
}

/**
   函数传递参数
 */
func adds(a, b int) int {
	return a + b
}
func subs(a, b int) int {
	return a - b
}
//传递函数
func functionValue(a, b int, do func(int, int) int) {
	fmt.Println(do(a, b))
}
func main() {
    a,b:=1,2
	fmt.Printf("原值 a:%v,b:%v \n", a, b)
     noChange(a,b)
	fmt.Printf("值传递后 a:%v,b:%v \n", a, b)
    fmt.Println("==========================")
	fmt.Printf("原值 a:%v,b:%v \n", a, b)
	noChangePoint(&a,&b)
	fmt.Printf("值传递后 a:%v,b:%v \n", a, b)

	functionValue(1, 1, adds)
	functionValue(1, 1, subs)




}