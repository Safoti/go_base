package main

import "fmt"

/**
   recover  恢复中断的程序
 */
func fullNamesss(first *string,last *string){
	defer recoverFullNames() //defer 激活
	if first ==nil{
		panic("runtime error: first name cannot be nil")
	}
	if last ==nil{
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *first, *last) //没有执行
	fmt.Println("returned normally from fullName")
}

func recoverFullNames() {
	if r := recover(); r!= nil {
		fmt.Println("recovered from ", r)
	}
}
func main() {
	/**
	  根据打印的结果可知.defer 都执行了
	 */
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullNamesss(&firstName, nil)
	fmt.Println("returned normally from main")


}
