package main

import "fmt"

/**
   panic 使用
 */
/**
   fullname 激活条件是，传递的参数不符合要求
    defer 延后执行
 */
func fullNamess(first *string,last *string){
	defer fmt.Println("deferred call in fullName")
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
//func slicePanic()  {
//	n:=[]int{5,7,4}
//	fmt.Println(n[4])//将产生错误
//	fmt.Println("normally returned from a")
//}
func main() {
	/**
	  根据打印的结果可知.defer 都执行了
	 */
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullNamess(&firstName, nil)
	fmt.Println("returned normally from main")

	//slicePanic()
	//fmt.Println("normally returned from main")

}
