package main

import "fmt"

func main() {
 var arr [4]int

 	//数组赋值
	for  i:=0;i<len(arr);i++ {
		arr[i] = i * 2
	}

	//遍历数组
	for i:=0; i < len(arr); i++ {
		fmt.Printf("Array at index %d is %d\n", i, arr[i])
	}
}
