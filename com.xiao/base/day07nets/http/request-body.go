package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/19
 * @Description   服务端请求体方式
 **/
func main() {
	mxu := http.NewServeMux()
	mxu.HandleFunc("/param", getRequestBody)
	ctx := context.Background()
	server := &http.Server{
		Addr:    ":9999",
		Handler: mxu,
		BaseContext: func(l net.Listener) context.Context {

			ctx = context.WithValue(ctx, "keyServerAddr", l.Addr().String())
			return ctx
		},
	}
	err := server.ListenAndServe()
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error listening for server: %s\n", err)
	}
}

/**
  处理body  参数
*/
func getRequestBody(w http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		fmt.Printf("could not read body: %s\n", err)
	}
	fmt.Println("读取的信息为：", string(body))

	io.WriteString(w, "This is my website!\n")
}

/**
  处理form 表单暂时不写了
*/
