package main

import "fmt"

type Tank interface {
	Tarea() float64
	Volume() float64

}
type myvalue struct {
	radius float64
	height float64
}

type Shaper interface {
	Area()float32
}
type Square struct {
	side float32
}
//函数
func (m myvalue)Tarea()float64  {
	return 2*m.radius*m.height +
		2*3.14*m.radius*m.radius
}//函数实现方法
func(m myvalue)Volume()float64{
	return 3.14 * m.radius * m.radius * m.height
}
//使用接口中的方法
func (sq *Square)Area()float32 {
	return sq.side * sq.side
}
func main() {
	var t Tank
	t=myvalue{10,14}
	fmt.Println("Area of tank :", t.Tarea())
	fmt.Println("Volume of tank:", t.Volume())

	sq1 := new(Square)
	sq1.side = 5

	var areaIntf Shaper
	areaIntf = sq1

	fmt.Printf("The square has area: %f\n", areaIntf.Area())


}
