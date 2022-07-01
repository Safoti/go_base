package main

import "fmt"

/**
 使用介绍
 */
func main() {
	//dir, err := ioutil.TempDir("","")
	//if err != nil {
	//	fmt.Errorf("failed to create temp dir: %v",err)
	//}
	//fmt.Println(dir)
    //if false {
    //   fmt.Println("没有错误")
	//}else {
	//	errors.New("测试错误的运行")
	//}

	//fmt.Println("hello")
	//for i := 0; i <=3; i++ {
	//	defer fmt.Println(i)
	//}
	//fmt.Println("world")


	//正确调用
	v,r:= div(100,2)
	if r!=nil{
		fmt.Println("(1)fail:",r)

	}else {
		fmt.Println("(1)succeed:",v)
	}

   //错误调用
	v,r=div(100,0)
	if r!=nil{
		fmt.Println("(2)fail:",r)
	}else {
		fmt.Println("(2)succeed:",v)
	}

}
// 自定义错误信息结构
type DIV_ERR struct {
	etype int  // 错误类型
	v1 int     // 记录下出错时的除数、被除数
	v2 int
}
//实现接口方法
func (div_err DIV_ERR) Error()string{
	if 0==div_err.etype {
		return  "除零错误"
	}else {
	  	return "其它未知错误"
	}
}
// 除法
func div(a,b int)(int ,*DIV_ERR){
	if b == 0 {
		// 返回错误信息
		return 0,&DIV_ERR{0,a,b}
	}else {
		// 返回正确的商
		return a / b, nil
	}
}