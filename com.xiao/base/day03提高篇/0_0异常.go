package main

import (
	"errors"
	"fmt"
	"net"
	"os"
	"path/filepath"
)
func main() {
  //one()

	//two()
	//three()
    four()
}
/**
  查找不存在的文件
*/
func one (){
	open, err := os.Open("/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(open.Name(),"打开txt 成功")
}
/**
    细化异常
 */
func two(){
	open, err := os.Open("test.txt")
	if err != nil {
	   if pErr,ok:=err.(*os.PathError);ok{
		   fmt.Println("没找到路径暂时无法发现", pErr.Path)
		   return
	   }
		fmt.Println("通用错误", err)
		return
	}
	fmt.Println(open.Name(),"打开txt 成功")
}
/**
      对应错误的处理方式
 */
func three()  {
	//不存在的地址
	add, err := net.LookupHost("Golang.ck")

	if err != nil {
		if dnsErr, ok :=err.(*net.DNSError);ok{
			if dnsErr.Timeout(){
				fmt.Println("操作超时")
				return
			}
			if dnsErr.Temporary(){
				fmt.Println("暂时错误")
				return
			}
			fmt.Println("通用的dns 错误", err)
			return
		}
		fmt.Println("通常错误", err)
		return
	}
	fmt.Println(add)
}

func  four()  {
	var ErrBadPattern=errors.New("syntax error in pattern")
	files, err := filepath.Glob("[")
	if err != nil {
		if err == ErrBadPattern {
			fmt.Println("Bad pattern error:", err)
			return
		}
		fmt.Println("Generic error:", err)
		return
	}
	fmt.Println("matched files", files)
}