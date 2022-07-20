package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/19
 * @Description   服务端  获取请求的参数
 **/
func main() {
	mxu := http.NewServeMux()
	mxu.HandleFunc("/param", getParamss)
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

func getParamss(w http.ResponseWriter, r *http.Request) {
	hasFirst := r.URL.Query().Has("first")
	first := r.URL.Query().Get("first")
	hasSecond := r.URL.Query().Has("second")
	second := r.URL.Query().Get("second")
	fmt.Printf(" got / request. first(%t)=%s, second(%t)=%s\n",
		hasFirst, first,
		hasSecond, second)
	io.WriteString(w, "This is my website!\n")

}
