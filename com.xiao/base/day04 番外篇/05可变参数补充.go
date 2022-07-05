package main

import "fmt"
/**
  知道可变参数的类型方式
 */
func  sum(t ...int)( res int)  {
	for _, v := range t {
	  	res+=v
	}
	return res
}
/**
  不知道可变参数的类型，但是可以预估就是几种类型的处理方式
 */
func sunNum(t ...interface{})(res float64){
	for _, tmp := range t {
		//条件判断数据类型
		switch v:=tmp.(type) {
		case int:
			res+=float64(v)
		case float64:
			res+=v
		case float32:
			res+=float64(v)
		}
	}
	return res
}
func main() {
  fmt.Println(sum(1,2,3,4,5))

	fmt.Println(sunNum(1,2.1,"asd",true))
}
