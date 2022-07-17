package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/17
 * @Description
 **/
func main() {
	val := []byte(`{"ID":1,"Name":"Reds","Colors":{"c":"Crimson","r":"Red","rb":"Ruby","m":"Maroon","tests":["tests_1","tests_2","tests_3","tests_4"]}}`)
    fmt.Println(jsoniter.Get(val,"Colors").ToString())
	fmt.Println(jsoniter.Get(val,"Colors","tests",0).ToString())
	//fmt.Println(jsoniter.Get(val, "colors", 0).ToString())//获取失败案例
}
