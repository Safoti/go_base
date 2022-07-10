package main

import (
	"fmt"
	"net"
	"strings"
)

/**
 * Created by safoti on 2022/7/9 21:54
**/
func main() {
	//1 创建服务端监听
	listen, err := net.Listen("tcp","0.0.0.0:8900")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()

	// 2.服务端不断等待请求处理

	for  {
		conn, err :=listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go ClientConn(conn)
	}
}
// 处理服务端逻辑
func ClientConn(conn net.Conn) {
	defer conn.Close()

	//获取客户端地址
	ipAddr:=conn.RemoteAddr().String()
	fmt.Println(ipAddr, "连接成功")
	buf:=make([]byte,1024)
	for  {
		n, err := 	conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}

	 result:=  buf[:n]
	 fmt.Printf("接收到数据，来自[%s]    [%d]:%s\n", ipAddr, n, string(result))
		// 接收到exit，退出连接
		if string(result) == "exit" {
			fmt.Println(ipAddr, "退出连接")
			return
		}
		// 回复客户端
		conn.Write([]byte(strings.ToUpper(string(result))))

	}
}
