package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/19
 * @Description   http   模拟的是服务器拦截根据请求的路径拦截响应的方法
 **/
func main() {
	//helloHttp()
}

//多路复用的http
func multableHttp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/s", getRoot)
	mux.HandleFunc("/s/hello", getHello)
	http.ListenAndServe(":9999", nil)

}

//hello world http

func helloHttp() {
	//需要处理的请求
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/hello", getHello)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		errors.New("失败了")
	}
}

/**
  http 请求
*/
func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}
func getHello(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /hello request\n")
	io.WriteString(w, "Hello, HTTP!\n")
}
