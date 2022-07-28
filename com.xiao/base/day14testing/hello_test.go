package main

import (
	"testing"
)
/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description
 **/
func TestHello(t *testing.T)  {
	got:=Hello()
	want:="Hello, world"
	if got !=want{
		t.Errorf("got %q want %q", got, want)
	}
	//运行测试命令 (指定文件)
	//go test -v hello_test.go  hello.go
}