package main

import (
	"errors"
	"fmt"
	"net/http"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/15
 * @Description    服务器端请求
 **/
func main() {
	http.HandleFunc("/", getRoots)
	err := http.ListenAndServe(":9999", nil)
	if err != nil {
		errors.New("失败了")
	}
}
func getRoots(w http.ResponseWriter, r *http.Request) {
	ms := r.Method
	hosts := r.URL.Host
	paths := r.URL.Path
	fmt.Println(r.URL.Port())
	fmt.Println(ms)
	fmt.Println(hosts)
	fmt.Println(paths)

	//fmt.Printf("got / request\n")
	//io.WriteString(w, "This is my website!\n")
}
