package main

import (
	"fmt"
	"sync"
)

func main() {
	//var mss map[string]string
	//mss=make(map[string]string)
	//fmt.Println(mss)
   m:=make(map[string]string)
   m["name"]="coding3"
   m["sex"]="man"
	for key := range m {
		fmt.Println("key:", key, ",value:", m[key])
	}

	delete(m,"name")

  //if value,ok:= m[key];ok{
  //  fmt.Println(key,value)
  //}else {
	//  fmt.Println(key, " 不存在")
  // }

	originalMap := make(map[string]int)
	originalMap["one"] = 1
	originalMap["two"] = 2

	targetMap := make(map[string]int)
	// Copy from the original map to the target map
	for key, value := range originalMap {
		targetMap[key] = value
	}
    //安全的map
    var scene sync.Map
	scene.Store("name","coding3min")
	scene.Store("age", 11)

   v,ok:=	scene.Load("name")
	if ok{
		fmt.Println(v)
	}
	v,ok=scene.Load("age")
	if ok{
		fmt.Println(v)
	}


	scene.Delete("age")

	scene.Range(func(key, value interface{}) bool {
		fmt.Println("key:",key,",value:",value)
		return true
	})

}
