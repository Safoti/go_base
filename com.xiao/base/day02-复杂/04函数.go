package main

import "fmt"

/**
  函数增强
 */
func main() {
	/**
	 go 的空白符，只想获取所需要的参数，别的可以放弃
	 */
	//原来调用
	mul, div := mul_div(105, 7)

	fmt.Println("105 x 7 = ", mul, div)
	//使用空白符
	mul, _ = mul_div(105, 7)
	fmt.Println("105 x 7 = ", mul)



	//defer 使用
	muls(23, 45)
	defer muls(23, 56) //defer 控制程序执行
	show()



	fmt.Println("Start")

	// Multiple defer statements
	// Executes in LIFO order
	defer fmt.Println("End")
	defer add(34, 56)
	defer add(10, 10)
}

func mul_div(n1 int, n2 int) (int, int) {

	// returning the values
	return n1 * n2, n1 / n2
}

// Functions
func muls(a1, a2 int) int {

	res := a1 * a2
	fmt.Println("Result: ", res)
	return 0
}

func show() {
	fmt.Println("Hello!, lktbz")
}
func add(a1, a2 int) int {
	res := a1 + a2
	fmt.Println("Result: ", res)
	return 0
}
