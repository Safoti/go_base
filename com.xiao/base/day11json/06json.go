package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
	"unsafe"
)
/**
 * @Author safoti
 * @Date Created in 2022/7/30
 * @Description   数据的读取
 **/

type City struct {
	Name string  `json:"name"`
	GDP int64      `json:"Gdp"`
	Popula int64   `json:"popula"`
	CreateDate time.Time `json:"create_date"`
	deleteDate *time.Time `json:"delete_date,-"` //-表示忽略的字段
	updateDate *time.Time`json:"update_date,omitempty"` //空值则忽略该字段 直接返回
}
type Country struct {
	Name string `json:"co_name"`
	Code string `json:"co_code"`
	City []City `json:"co_city"`
}

func main() {
	//fs, err := os.Open("go.json")
	//if err !=nil{
	//	panic(err)
	//}
	var cos Country
	re,err:=ioutil.ReadFile("go.json")
	if err !=nil{
		panic(err)
	}
	json.Unmarshal(re,&cos)
	fmt.Println(cos)
}
