package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/24
 * @Description 
 **/
func main() {
		var notes[7]string
		notes[0]="do"
		notes[1]="re"
		notes[2]="mi"
		fmt.Println(notes[0])
		fmt.Println(notes[1])

		var parmes[5]int
		parmes[0]=2
		parmes[2]=4
		fmt.Println(parmes[0])
		fmt.Println(parmes[2])
		fmt.Println(parmes[4])//打印0值

		//字符串的默认值
		var nos[4]string
		nos[0]="do"
		fmt.Println(nos[0])
		fmt.Println("--",nos[2])//打印结果
	var ps=	new([4]int)
	ps[0]=2
	ps[1]=3
	fmt.Println(ps[1])
	fmt.Println(ps[2]) //new 方式创建也是0值


	noss:=[7]string{"do","re","mi","fa","sa"}
	fmt.Println(noss[2])


	text:=[7]string{
		"this",
		"is",
		"test",
	}
	for i, s := range text {
		fmt.Println(i,"==",s)
	}
}
