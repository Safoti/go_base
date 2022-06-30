package main

import "fmt"

func main() {
	x := getMin(1, 3, 2, 0)
	fmt.Printf("The minimum is: %d\n", x)

	slice := []int{7,9,3,5,1}

	x = getMin(slice...)
	fmt.Printf("The minimum in the slice is: %d", x)

	a()
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