package main

import (
	"fmt"
	"reflect"
)

/**
 * Created by safoti on 2022/7/9 18:46
**/
type ReflectDemo struct {
    Name string
    Method string
    Vars string
    class []ClassDemo
}
type ClassDemo struct {
	 CName string
	 CPath string
}

type common struct {
	users Users
}
type Users struct {
	Name string
	Age int
	Address string
}

func (u *Users)getName(name string)  {
	fmt.Println(u.Name==name)
}
func (u Users)getNames(name string)  {
	fmt.Println(u.Name==name)
}
func (c *common)getName(us Users)  {
	fmt.Println(us.Name)
}

func (rs *ReflectDemo)getFredls(name string)  {

}

func (rs ReflectDemo)get(name string){

}
func main() {
   //tandfy:=reflect.TypeOf("andy")
	//tandfy1:=reflect.TypeOf(1)
	//fmt.Println(tandfy)
	//fmt.Println(tandfy1)
	cd:=[]ClassDemo{
		{	CName: "数学",
			CPath: "3年2班"},{
			CName: "语文",
			CPath: "3年1班"},{
			CName: "外文",
			CPath: "3年3班"},
	}
	red:=ReflectDemo{Name: "zs",Method: "add",Vars: "go",class:cd,}

	res:=reflect.TypeOf(&red)
    //fmt.Println(res)  //main.ReflectDemo
	//fmt.Println(res.Name()) //ReflectDemo  结构体名称
	//fmt.Println(res.Kind()) //struct


	//fileNums:=res.NumField()
	//for i := 0; i < fileNums; i++ {
	//	//获取结构体中的属性
	//	filed:=res.Field(i)
	//	fmt.Println(filed.Name)
	//	fmt.Println(filed.Type)
	//	fmt.Println(filed.Index)
	//	fmt.Println(filed.Tag)
	//	fmt.Println(filed.Anonymous)
	//	fmt.Println(filed.PkgPath)
	//}
	//fmt.Println(" ===============")
	//这两两种也是获取结构体中的属性
	//if namefield,ok:= res.FieldByName("Name");ok{
	//   	  fmt.Println(namefield.IsExported())
	//   }
	//   namef:=res.FieldByIndex([]int{0})
    //  fmt.Println(namef.Name)


	//获取结构体中的方法
	//methodNum :=res.NumMethod()
	//for i := 0; i < methodNum; i++ {
	//	method := res.Method(i)
	//	fmt.Printf("method name:%s ,type:%s, exported:%t\n", method.Name, method.Type, method.IsExported())
	//
	//}
	//fmt.Println()

}
