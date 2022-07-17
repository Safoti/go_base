package main

import "github.com/kataras/iris/v12"

/**
 * @Author safoti
 * @Date Created in 2022/7/15
 * @Description   模拟一个服务方便接收与返回信息
 **/
func main() {
    app:= iris.New()
    app.Post("/test", func(ctx iris.Context) {
		  ctx.WriteString("我是服务端，我正在返回信息")
	})
    app.Listen(":8090")
}
