package main

import (
	"fmt"
	"sync"
)

/**
 * Created by safoti on 2022/7/11 9:36
**/
func main() {

	var wg sync.WaitGroup
	sayWgHello:= func() {
      defer wg.Done()
      fmt.Println("hello")
	}
	wg.Add(1)
	go sayWgHello()
	wg.Wait()


	//修改值
	slaution:="hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		slaution= "welcome"
	}()
	wg.Wait()
	fmt.Println(slaution) //welcome



	//继续

for _,gss := range[ ]string{"hs", "js", "gs"}{
   wg.Add(1)
   go func() {
   	defer wg.Done()
   	fmt.Println(gss) //gss gss gss
   }()
}
wg.Wait()



//
}
