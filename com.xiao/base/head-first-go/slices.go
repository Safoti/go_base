package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/24
 * @Description  切片与数组最最主要的区别是：没有容量
 **/
func main() {
   ins:= []int{2,3,4}
	fmt.Println(ins)


   array1:=[5]string{"a","b","c","d","e"}
   slice1:=array1[0:3]
   array1[0]="x" //修改数据
   fmt.Println(array1)
   fmt.Println(slice1)

	fmt.Println("====================")
	array2:=[5]string{"h","i","j","k","l"}
	slice2:=array1[2:5]
	fmt.Println(slice2)
    slice2[1]="x"
   	fmt.Println(array2)
   	fmt.Println(slice2)

	array3:=[5]string{"h","i","j","k","l"}
	slice3:=array3[0:3]
	slice4:=array3[2:5]
	fmt.Println(slice3)
	array3[2]="x"
	fmt.Println(array3)
	fmt.Println(slice3,slice4)
}
