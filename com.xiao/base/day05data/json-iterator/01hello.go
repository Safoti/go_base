package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)
type ColorGroup struct {
	ID int
	Name   string
	Colors []string
}


func main() {
	//构建数据
	group := ColorGroup{
		ID:     1,
		Name:   "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	//转换成json
	b, err := jsoniter.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	//打印转换的json
	fmt.Println(string(b))

	var jsonBlob = []byte(`[{"Name": "Platypus", "Order": "Monotremata"}, 
{"Name": "Quoll",    "Order": "Dasyuromorphia"} ]`)
	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	es := jsoniter.Unmarshal(jsonBlob, &animals)
	if es != nil {
		fmt.Println("error:", es)
	}
	fmt.Printf("%+v", animals)

	val := []byte(`{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}`)
	s:=jsoniter.Get(val,"Colors",0).ToString()
	fmt.Println(s)

}
