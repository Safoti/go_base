package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/11
 * @Description
 **/
func main() {
     insd:=make(chan int)
     go func() {
     	defer close(insd)
		 for i := 0; i <=5 ; i++ {
			 insd<-i
		 }
	 }()
	for  interger := range insd {
		fmt.Printf("%v ",interger)
	}
}
