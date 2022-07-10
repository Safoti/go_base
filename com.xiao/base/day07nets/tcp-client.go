package main

import (
	"fmt"
	"net"
)

/**
 * Created by safoti on 2022/7/9 21:58
**/
func main() {
	conn, err := net.Dial("tcp","0.0.0.0:8900")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	// 缓冲区

	for  {
		buf := make([]byte, 1024)
		fmt.Printf("请输入发送的内容：")
		fmt.Scan(&buf)
		fmt.Printf("发送的内容：%s\n", string(buf))
		// 发送数据
		conn.Write(buf)

		res := make([]byte, 1024)
		n, err := conn.Read(res)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := res[:n]
		fmt.Printf("接收到数据:%s\n", string(result))

	}
}
