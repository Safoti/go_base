package main

import (
	"encoding/json"
	"fmt"
)

/**
 json 中嵌套对象
 */

type JsonSingle struct {
	Name string
}
type Jsons struct {
	js []JsonSingle
}
type  Java struct {
	Name string
	Level int
	UsePeson string
	json Jsons `json:"jsons"`
}
func main() {
	js:=make([]JsonSingle,5)
	js[0]=JsonSingle{"JAVA"}
	js[1]=JsonSingle{"golang"}
	js[2]=JsonSingle{"python"}
	JSS:=Jsons{js: js}
	jsc:=&Java{
		Name: "TEST",
		Level: 2,
		UsePeson: "ALL",
		json:JSS,
	}

	marshal, err := json.Marshal(jsc)
	if err != nil {
		return
	}
	fmt.Println(string(marshal))



}
