package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/11 12:33
 * @Description
 **/
func main() {
	//三种流创建方式   输出流  输入流  双向流
	//var rec<-chan interface{} 输出流
	//var send chan <-interface{} 输入流
    //datas:=make(chan interface{}) 双向流

    //定义int类型的流
	//intStream:=make(chan int)


	//简单的例子

	stringStream:=make(chan string)
	go func() {
		stringStream<-"hello world channel" //流写入管道
	}()
	//fmt.Println(<-stringStream) //从管道获取流

	//读取两个值
	slatation,ok:=<-stringStream
	fmt.Printf ("( %v): %v ", ok, slatation)



	insts:=make(chan  int)
	close(insts)
	ist,ok:=<-insts
	fmt.Printf(" %v : %v ",ok,ist)
}
