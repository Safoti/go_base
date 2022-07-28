package main

import (
	"fmt"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/26
 * @Description   生产者模式演示
 **/
func main() {
	//Producer(1) //顺序执行
	//Producer(2)

	go Producer(2)
	go Producer(3)

	//匿名函数
	go func() {
		for i := 0; i <5 ; i++ {
			fmt.Printf("Foo message %v\n",i)
		}
		//继续调用函数
		Producer(4)
	}()
	time.Sleep(5*time.Millisecond) //暂时使用睡眠程序
}

func Producer(is int) {
	for i := 0; i <5 ; i++ {
		fmt.Printf("工厂代号：%d 生产了第%d 个产品\n",is,i)
	}


}
