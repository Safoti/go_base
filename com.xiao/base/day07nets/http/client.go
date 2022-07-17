package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/15
 * @Description  模拟客户端发送的请求
 **/
func main() {
     // sendGet()
	//sendPost()
}


/**

   post 请求
 */
func sendPost()  {
	res, err := http.Post("/test", "none", nil)
	if err!=nil{
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", body)
}
/**
     get 请求
 */
func sendGet(){
	//官网demo

	res, err := http.Get("http://www.baidu.com")
	/**
	  get 方法点进去 发现  默认调用 http 下的 client get 方法    return DefaultClient.Get(url)
      //接着是构建请求
	   req, err := NewRequest("GET", url, nil)




	*/

	if err != nil {     log.Fatal(err) }
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body) }
	if err != nil {     log.Fatal(err) }
	fmt.Printf("%s", body)
}
