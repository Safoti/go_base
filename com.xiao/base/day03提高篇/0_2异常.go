package main

import (
	"errors"
	"fmt"
	"math"
)

/**
   使用官方   将错误信息返回给调用者
 */
func circleArea(radius float64)(float64,error){
	if radius <0{
		//官方自定义异常
		return 0,errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi*radius*radius,nil
}
/**
  追加更多的错误信息
 */
func circleArea2(radius float64)(float64,error){
	if radius < 0 {
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}
func main() {
  radius:=-20.2
  //radius:=20.2
	//area, err := circleArea(radius)
	area, err := circleArea2(radius)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Area of circle %0.2f", area)
}
