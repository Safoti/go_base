package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	newRead()
}
func newRead(){
	reader := 	bufio.NewReader(strings.NewReader("http://studygolang.com. \nIt is the home of gophers"))
	line,_:=reader.ReadBytes('\n')
	fmt.Printf("the line:%s\n", line)
	n, _ := reader.ReadBytes('\n')
	fmt.Printf("the line:%s\n", line)
	fmt.Println(string(n))
}