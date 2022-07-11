package main

import "fmt"

func main() {
	var mapLit map[string]int  //创建方式一
	mapLit=map[string]int{"one":1,"two":2} //创建方式二
  	fmt.Println(mapLit)
	m:=make(map[string]int)//创建方式三
	m["k1"]=7
	m["k2"]=13
	fmt.Println("map:",m)

	v1:=m["k1"]
   fmt.Println("v1:",v1)

	fmt.Println("len():",len(m))

	delete(m,"k2")
	fmt.Println("map",m)

    _, prs:=m["k2"]
	fmt.Println("prs:", prs)

	var value int
	var isPresent bool
    map1:=make(map[string]int)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25
	value, isPresent =map1["Beijing"]
	if isPresent {
		fmt.Printf("The value of \"Beijing\" in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Beijing")
	}
	value, isPresent = map1["Paris"]
	fmt.Printf("Is \"Paris\" in map1 ?: %t\n", isPresent)
	fmt.Printf("Value is: %d\n", value)
	map12 := make(map[int]float32)
	map12[1] = 1.0
	map12[2] = 2.0
	map12[3] = 3.0
	map12[4] = 4.0
	for i, v := range map12 {
		fmt.Printf("key is: %d - value is: %f\n", i, v)
	}
	capitals := map[string] string {"France":"Paris", "Italy":"Rome", "Japan":"Tokyo" }
	for skey  := range capitals {
		fmt.Println("Map item: Capital of", skey, "is", capitals[skey])
	}
	for _ ,value := range capitals {
		fmt.Println("Map item: Capital value", value)
	}


	employeeSalary := 	map[string]int{
		"steve": 12000,
		"jamie": 15000,
		"mike": 9000,
	}
	employeeSalary["gogo"]=9999
	newEmp:="jamie"
	fmt.Println(employeeSalary)
	/**
	value, ok := employeeSalary[newEmp]
	    if ok == true {
	        fmt.Println("Salary of", newEmp, "is", value)
	        return
	    }
	 使用下边的进行替换
	 */
	if vals,ok:=employeeSalary[newEmp];ok{
		fmt.Println("Salary of", newEmp, "is", vals)
		return
	}
	fmt.Println(newEmp,"not found")


	//元素的删除
	delete(employeeSalary, "steve")
	fmt.Println("map after deletion", employeeSalary)
}
