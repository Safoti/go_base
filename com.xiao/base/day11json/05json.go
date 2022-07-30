package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
)
/**
 * @Author safoti
 * @Date Created in 2022/7/30
 * @Description   格式化显示的数据进行修改
 **/

type City struct {
	Name string  `json:"name"`
	GDP int64      `json:"gdp"`
	Popula int64   `json:"popula"`
	CreateDate customTime `json:"create_date"`
	deleteDate customTime `json:"delete_date,-"` //-表示忽略的字段
	updateDate customTime `json:"update_date,omitempty"` //空值则忽略该字段 直接返回
}
type Country struct {
	Name string `json:"co_name"`
	Code string `json:"co_code"`
	City []City `json:"co_city"`
}
type customTime struct {
	 time.Time
}
const  layout ="2006-01-02"
//自定义格式
func (c customTime)MarshalJSON()([]byte,error){
	return []byte(fmt.Sprintf("\"%s\"",c.Format(layout))),nil
}
func (c *customTime)UnmarshalJSON(v []byte)(error)  {
	var err error
	c.Time,err=time.Parse(layout,strings.ReplaceAll(string(v),"\n",""))
	if err!=nil{
		return err
	}
	return  nil
}
func main() {
	//构建数据
	co:=Country{
		Name: "CN",
		Code: "CN1000",
		City: []City{
			{Name: "BJ",GDP: 10000,Popula: 130000,CreateDate:customTime{ time.Now()}},
			{Name: "SH",GDP: 12000,Popula: 150000,CreateDate: customTime{time.Now()}},
			{Name: "GZ",GDP: 9000,Popula: 11000,CreateDate: customTime{time.Now()}},
		},
	}
	//拼装json 并格式化json
	b, err := json.Marshal(co)
	if err !=nil{
		panic(err)
	}
	fmt.Println(string(b))
	cos:=Country{}
	//f,err:=os.Open("go.json")
	re,err:=ioutil.ReadFile("go.json")
	if err !=nil{
		panic(err)
	}

	json.Unmarshal(re,&cos)
	fmt.Printf("%+v\n",cos)
}


