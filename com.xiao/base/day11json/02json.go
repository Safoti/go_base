package main

import (
	"encoding/json"
	"fmt"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description   字符串转换成json或者json 转换成字符串
 **/
func main() {
	var jsonBlob = []byte(`[
	{"Name": "Platypus", "Order": "Monotremata"},
	{"Name": "Quoll",    "Order": "Dasyuromorphia"}
]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob,&animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", animals)
}

