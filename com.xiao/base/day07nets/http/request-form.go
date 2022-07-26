package main

import (
	"fmt"
	"net/http"
	"log"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/19
 * @Description   服务端  获取form 表单数据  表单数据通过postman 进行模拟
 **/
func main() {
   http.HandleFunc("/",getForm)
	err := http.ListenAndServe("localhost:9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	fmt.Println("服务启动成功")
}

func getForm(writer http.ResponseWriter, request *http.Request) {
	request.ParseForm()
	for k, v := range request.Form {
		fmt.Println("key is: ", k)
		fmt.Println("val is: ", v)
	}
	fmt.Fprintf(writer, "hello world")
}

