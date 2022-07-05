package main

import "fmt"
/**
  闭包
 */
func shareDataSample()func(){
	count:=0
	return func() {
		count++
		fmt.Printf("调用次数 %v \n", count)
	}
}
func main() {
c1,c2:=shareDataSample(),shareDataSample()
c1()
c2()
c1()
c2()
c1()
}
