package main

import (
	//"bytes"
	"encoding/json"
	"fmt"
	"os"

	//"os"
	"time"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/30
 * @Description    结构体字段调整
 **/
type City struct {
	Name string  `json:"name"`
	GDP int64      `json:"gdp"`
	Popula int64   `json:"popula"`
	CreateDate time.Time `json:"create_date"`
	deleteDate *time.Time `json:"delete_date,-"` //-表示忽略的字段
	updateDate *time.Time `json:"update_date,omitempty"` //空值则忽略该字段 直接返回
}
type Country struct {
	Name string `json:"co_name"`
	Code string `json:"co_code"`
	City []City `json:"co_city"`
}

func main() {
	//构建数据
	times :=time.Now()

	co:=Country{
		Name: "CN",
		Code: "CN1000",
		City: []City{
			{Name: "BJ",GDP: 10000,Popula: 130000,CreateDate:  times},
			{Name: "SH",GDP: 12000,Popula: 150000,CreateDate: times},
			{Name: "GZ",GDP: 9000,Popula: 11000,CreateDate: times},
		},
	}


	//b, err := json.Marshal(co)
	//if err !=nil{
	//	panic(err)
	//}
	////格式化
	//var out bytes.Buffer
	//json.Indent(&out,b,"","\t")
	//out.WriteTo(os.Stdout)

	//拼装json 并格式化json
	b, err := json.MarshalIndent(co, "", "\t")
	if err !=nil{
		panic(err)
	}
	fmt.Println(string(b))
	os.WriteFile("go.json",b,0666)
}
/**
问题  日期显示方式需要修改  需要使用中间态进行修改
{"co_name":"CN","co_code":"CN1000","co_city":[{"name":"BJ","Gdp":10000,"popula":130000,"create_date":"2022-07-30T16:54:3
8.9833406+08:00"},{"name":"SH","Gdp":12000,"popula":150000,"create_date":"2022-07-30T16:54:38.9833406+08:00"},{"name":"G
Z","Gdp":9000,"popula":11000,"create_date":"2022-07-30T16:54:38.9833406+08:00"}]}
 */

