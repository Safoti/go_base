package main

import "fmt"

func main() {
  var a chan int  //传递int 型数据
  if a ==nil{
	  fmt.Println("channel a is nil, going to define it")
   a= make(chan int)
	  fmt.Printf("Type of a is %T", a)

   //data:=<-a read from channel a
   //a<-data   write to channel a
  }
}
