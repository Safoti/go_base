package main

import (
	"fmt"
	"math"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/17
 * @Description
 **/

type Point struct {
	X,Y float64
}
//普通函数
func Distance(p,q Point)float64  {
	return math.Hypot(p.X-p.X,q.Y-p.Y)
}
//方法
func (p Point)Distance(q Point)float64 {
	return math.Hypot(p.X-p.X,q.Y-p.Y)
}

func main() {
	p:=Point{1,2}
	q:=Point{4,6}
	fmt.Println(Distance(p,q))//函数调用
	fmt.Println(p.Distance(q)) //方法调用
}
