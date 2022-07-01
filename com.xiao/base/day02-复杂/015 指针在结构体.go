package main

import "fmt"

type Employees struct {

	// taking variables
	name  string
	empid int
}

func main() {
	emp := Employees{"ABC", 19078}
	pts:=&emp
	println(pts)
	println(pts.name)
	println(pts.empid)
	println((*pts).name)



	//修改属性
	pfs:=&emp
	pfs.name= "XYZ"
	fmt.Println(pfs.name)
}
