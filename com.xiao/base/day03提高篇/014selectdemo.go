package main

import (
	"fmt"
	"time"
)

/**
将上述程序中的函数命名为server1and的原因server2是为了说明select的
实际使用。

假设我们有一个关键任务应用程序，我们需要尽快将输出返回给用户。
此应用程序的数据库被复制并存储在世界各地的不同服务器中。
假设函数server1和server2实际上正在与 2 个这样的服务器进行通信。
每个服务器的响应时间取决于每个服务器的负载和网络延迟。
select我们将请求发送到两个服务器，然后使用该语句在相应的通道上等待响应。
首先响应的服务器由 select 选择，而另一个响应被忽略。
这样我们就可以将相同的请求发送到多个服务器，并将最快的响应返回给用户
 */
func  server1(ch chan string)  {
	time.Sleep(6000*time.Millisecond)
    ch <-"from server1"
}

func  server2 (ch chan string)  {
	time.Sleep(3000*time.Millisecond)
	ch <-"from server2"
}
func main() {
	out1:=make(chan string)
    out2:=make(chan string)
    go server1(out1)
    go server2(out2)
	/**
	 select  会产生阻塞 ，等待数据准备完毕
	 */
	/**
    select {
	 case s1:=<-out1:
		fmt.Println(s1)
	 case s2:=<-out2:
		fmt.Println(s2)

	}**/

	/**
	使用循环进行改写
	 */
	for  {
		time.Sleep(1000*time.Millisecond)
		select {
		case s1:=<-out1:
			fmt.Println(s1)
		case s2:=<-out2:
			fmt.Println(s2)
		default:
			fmt.Println("  data is not ready ")
		}
	}
}
