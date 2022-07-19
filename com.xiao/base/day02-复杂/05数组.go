package main

import "fmt"

func main() {
 var arr [4]int
 	fmt.Println(arr)
	fmt.Println(arr[0])
 	//数组赋值
	for  i:=0;i<len(arr);i++ {
		arr[i] = i * 2
	}

	//遍历数组
	for i:=0; i < len(arr); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr[i])
	}

	var q [3]int=[3]int{1,2,3}
	fmt.Println(q)
	//有元素的个数决定数组的长度
	p:=[...]int{10,121,151,51}
	fmt.Println(p)
	fmt.Println(len(p))
}
