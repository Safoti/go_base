package main

import "fmt"

//缓冲区的channel
func main() {
 ch:=make(chan  string ,2) //2个缓冲区
  ch<-"navi"
  ch<-"4am"
  fmt.Println(<-ch)
  fmt.Println(<-ch)
}
