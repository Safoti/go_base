package main

import "fmt"

type People struct {
	name string
}
func (p *People)toString(){
	fmt.Println(p.name)
	fmt.Printf("p的地址 %p \n", &p)
}
func main() {
   p1:=People{"coding3min"}
   p1.toString()
   //声明时赋值
   //body2:=Body{
   //	"tom",13
   //}

	//bodys := []Body{
	//   Body{"jack", 12}, Body{"lynn", 18},
   //}

	//class1 := struct {
	//	bodys []Body
	//}{
	//	[]Body{Body{"jerry", 24}},
	//}


}
