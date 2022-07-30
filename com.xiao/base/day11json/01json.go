package main

import (
	"encoding/json"
	"fmt"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description   标准库之json 之结构体转换成json
   字符串转换成json或者json 转换成字符串
 **/
func main() {
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
	}
	group:=ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
}
