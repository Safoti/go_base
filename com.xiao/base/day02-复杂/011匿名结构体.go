package main

import "fmt"
//匿名字段
type Stus struct {
	int
	string
	float64
}
func main() {
	//匿名结构体
	Ele:= struct {
		name,branch,language string
		Particles int
	}{
		name:      "Pikachu",
		branch:    "ECE",
		language:  "C++",
		Particles: 498,
	}
	fmt.Println(Ele)
	//调用匿名字段
	values:=Stus{123,"lk",666.66}
	//直接通过数据类型进行了调用
    fmt.Println("匿名字段 int:",values.int)
    fmt.Println("匿名字段 string:",values.string)
    fmt.Println("匿名字段 float64:",values.float64)
}

