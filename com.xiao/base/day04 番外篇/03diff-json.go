package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name    string
	Age     int
	Roles   []string
	Skill   map[string]float64
}
//一个json 中有复杂的数据类型
func main() {
 	skil:=make(map[string]float64)
	skil["python"] = 99.5
	skil["elixir"] = 90
	skil["ruby"] = 80.0
    ros:=[]string{"own","admin","player"}

  user:=  User{
    	Name: "lktbz",
    	Age: 10,
    	Roles: ros,
    	Skill: skil,
    	}
	marshal, err := json.Marshal(user)
	if err != nil {
		panic(err)
		return
	}
	fmt.Println(string(marshal))

	stsL:="{\n    \"Name\": \"lktbz\", \n    \"Age\": 10, \n    \"Roles\": [\n        \"own\", \n        \"admin\", \n        \"player\"\n    ], \n    \"Skill\": {\n        \"elixir\": 90, \n        \"python\": 99.5, \n        \"ruby\": 80\n    }\n}"
	us:=&User{}
	json.Unmarshal([]byte(stsL),us)
	fmt.Println(us)
}
