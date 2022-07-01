package main

import "fmt"

func main() {
	x := getMin(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)

	slice := []int{7,9,3,5,1}

	x = getMin(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)

	a()

	//匿名函数
	value:= func() {
		fmt.Println("测试 匿名函数")
	}
	value()

	//匿名函数传递参数  并且自调用
	func(ele string){
		fmt.Println(ele)
	}("lktbz")


	//函数当参数传递调用
	vas :=func(p,q string)string{  //自定义函数
		return  p+q+"LK"
	}
	LKL(vas)
	//函数当返回值使用
	sd := GFG()
	fmt.Println(sd("Welcome ", "to "))
}
/**
  传递的可变参数比较大小，找出最小的数并返回
 */

func getMin(s ...int)(int){
	if len(s)==0{
		return 0
	}
	min:=s[0]
	for _, v := range s {
      if v<min{
      	min=v
	  }
	}
	return  min
}

func a() {
	i := 0
	defer fmt.Println(i)
	i++
	fmt.Println(i)
	return
}
//函数当做参数传递
//传递匿名函数， 函数传入p,q 字符串并且返回一个字符串
func LKL(i func(p,q string)string){
	fmt.Println(i ("lks", "for"))
}

//函数返回另一个函数
//                   返回值类型 func(i,j string)string
func  GFG()func(i,j string)string  {
	//运算会返回的值
	myf:= func(i,j string)string {
		 return i + j + "lktbz"
	}
	//返回
	return myf
}

