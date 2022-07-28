package main

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description
 **/

type(
	/**
	   接口嵌套
	 */
	Shape interface {
		Area() float64
	}
	Circle interface {
		Radius() float64
		Shape
	}
	Rect interface {
		Width() float64
		Length() float64
		Shape
	}
)
func main() {


}
