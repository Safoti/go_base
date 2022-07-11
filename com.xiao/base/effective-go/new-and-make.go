package main

import "fmt"

/**
 * Created by safoti on 2022/7/10 21:30
Remember that make applies only to maps, slices and channels
and does not return a pointer. To obtain an explicit pointer
allocate with new or take the address of a variable explicitly.

**/
func main() {
	var p *[]int = new([]int)
   fmt.Println(p)
	//var v  []int = make([]int, 100)
   //fmt.Println(v)

	for i := 0; i < 100; i++ {
		*p= append(*p,i)
	}
	fmt.Println(p)


	isns:=make([]int,10)
	for i := 0; i <10 ; i++ {
		isns[i]=i
	}
	//fmt.Println(isns)
    fmt.Println("[0:] ",isns[0:])
    fmt.Println("[1:] ",isns[1:])
    fmt.Println("[:9] ",isns[:9])
    fmt.Println("[2:5] ",isns[2:5])
    fmt.Println("[3:4] ",isns[3:4])
    fmt.Println("[:] ",isns[:])


	hw:=append([]byte("Hello"),"World"...)
    fmt.Println(string(hw))



}
