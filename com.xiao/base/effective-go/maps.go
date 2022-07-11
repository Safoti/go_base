package main

import "fmt"

/**
 * Created by safoti on 2022/7/10 22:17
**/
func main() {
   var timeZone=map[string]int{
	   "UTC":  0*60*60,
	   "EST": -5*60*60,
	   "CST": -6*60*60,
	   "MST": -7*60*60,
	   "PST": -8*60*60,
   }

     offset:= timeZone["EST"]
	 fmt.Println(offset)
      //通过key 判断value 是否存在
      if vas,ok:=timeZone["GO"];ok{
       	 fmt.Println(vas)
	  }else {
	  	fmt.Println("值不存在....")
	  }



     // 删除  delete(timeZone,"PST")

}
