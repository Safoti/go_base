package main

import "fmt"

/**
   panic 使用
 */
/**
   fullname 激活条件是，传递的参数不符合要求
 */
func fullName(first *string,last *string){
	if first ==nil{
		panic("runtime error: first name cannot be nil")
	}
	if last ==nil{
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *first, *last)
	fmt.Println("returned normally from fullName")
}


/**
     程序本身产生的错误
 */
func slicePanic()  {
	n:=[]int{5,7,4}
	fmt.Println(n[4])//将产生错误
	fmt.Println("normally returned from a")
}
func main() {
	/**
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
	**/
	slicePanic()
	fmt.Println("normally returned from main")





}
