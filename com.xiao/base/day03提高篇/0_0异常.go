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
    //four()
	//errorIS()
	errorAs()
	//errorUnwrap()
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
		//异常进行分类处理
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
			//分类细化错误，对其进行处理
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
	// New returns an error that formats as the given text.
	// Each call to New returns a distinct error value even if the text is identical.
	var ErrBadPattern=errors.New("syntax error in pattern")
	//fmt.Println(ErrBadPattern)
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

type ErrorString struct {
	s string
}

func (e *ErrorString) Error() string {
	return e.s
}
/**

error AS  提取指定的错误类型
*/
func  errorAs() {
	var targetErr *ErrorString
	err := fmt.Errorf("new error:[%w]", &ErrorString{s:"target err"})
	fmt.Println(errors.As(err, &targetErr))

}
/**

  判断被包装的错误信息是否包含指定错误
 */
func errorIS()  {
	err1 := errors.New("new error")
	err2 := fmt.Errorf("err2: [%w]", err1)
	err3 := fmt.Errorf("err3: [%w]", err2)

	fmt.Println(errors.Is(err3, err2))
	fmt.Println(errors.Is(err3, err1))
}
/**
   将包装的错误拆出来
 */
func errorUnwrap()  {
	err1:=errors.New("我是真正的报错信息")
	err2 := fmt.Errorf("err:[%w]",err1)
	err3 := fmt.Errorf("err3: [%w]", err2)
	fmt.Println(err3)
	fmt.Println(errors.Unwrap(err3))
	fmt.Println(errors.Unwrap(errors.Unwrap(err3)))


}