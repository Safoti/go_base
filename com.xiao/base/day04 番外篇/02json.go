package main

import (
	"encoding/json"
	"fmt"
)

// Product 商品信息
//type Product struct {
//	Name      string
//	ProductID int64
//	Number    int
//	Price     float64
//	IsOnSale  bool
//}
//type Product struct {
//	Name string `json:"name"`
//	ProductID int64 `json:"-"` //表示不进行序列化
//	Number int `json:"number"`
//	Price float64 `json:"Price"`
//	IsOnSale bool `json:"is_on_sale,string"`
//}
//type Product struct {
//	Name      string  `json:"name"`
//	ProductID int64   `json:"product_id,omitempty"`   //tag里面加上omitempy，可以在序列化的时候忽略0值或者空值
//	Number    int     `json:"number"`
//	Price     float64 `json:"price"`
//	IsOnSale  bool    `json:"is_on_sale,omitempty"`
//}

/**
  需要序列化的字段与结构体字段不一致
 */
type Product struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,string"`
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

func main() {
  p:=&Product{}
	p.Name = "iphone 13 pro"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 0
	//转换成json
	data, err := json.Marshal(p)
	if err != nil {
		return
	}
	fmt.Println(string(data))




	str := `{"name":"test","product_id":"1","number":"110011","price":"0.01","is_on_sale":"true"}`
	ps:=Product{}
	//json to object
	json.Unmarshal([]byte(str),&ps)
	fmt.Println(ps)



}
