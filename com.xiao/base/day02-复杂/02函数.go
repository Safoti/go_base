package main

import "fmt"

var num int = 10
var numx2, numx3 int
var min, max int
func main() {
	numx2, numx3 = getX2AndX3(num)
	PrintValues()

	numx2, numx3 = getX2AndX3_2(num)
	PrintValues()


	var i1 int
	var f1 float32
	i1, _, f1 = freeValue()
	fmt.Printf("The int: %d, the float: %f \n", i1, f1)


	min, max = minMax(78, 65)
	fmt.Printf("Minmium is: %d, Maximum is: %d\n", min, max)
}

func PrintValues() {
	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
}
//默认返回
func getX2AndX3(input int)(int ,int){
	return 2*input,3*input

}
//返回具体的参数
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2=2 * input
	x3=3 * input
	return
}

/**
   空白符返回
 */

func  freeValue() (int ,int ,float32) {
	return 5, 6, 7.5
}
/**
   比较大小
 */

func minMax(a int ,b int)(min int,max int){
	if a<b{
		min=a
		max=b
	}else {
		max=a
		min=b
	}
	return
}