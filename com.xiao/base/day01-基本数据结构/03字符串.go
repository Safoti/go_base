package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is an example of a string"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s? ", str, "Th")
	fmt.Println("=======前缀测试========")
	fmt.Printf("%t\n",strings.HasPrefix(str,"Th"))

	var strs string = "Hi, I'm Marc, Hi."
	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Println("=======元素出现的下标========")
	fmt.Printf("%d\n", strings.Index(strs, "Marc"))
	fmt.Printf("%d\n", strings.Index(strs, "Hi"))
	fmt.Printf("%d\n", strings.LastIndex(strs, "Hi"))
	fmt.Printf("%d\n", strings.Index(strs, "Burger"))
	var strss string = "Hello, how is it going, Hugo?"
	var manyG = "gggggggggg"

	fmt.Println("=======元素出现的个数========")
	fmt.Printf("%d\n", strings.Count(strss, "H"))
	fmt.Printf("%d\n", strings.Count(manyG, "gg"))
	var origS string = "Hi there! "
	var newS string
	fmt.Println("=======复制字符串========")
	newS  =strings.Repeat(origS,3)
	fmt.Printf("The new repeated string is: %s\n", newS)

	var orig string = "Hey, how are you George?"
	var lower string
	fmt.Println("=======大小写字符串========")
	var upper string
	fmt.Printf("The original string is: %s\n", orig)
	lower = strings.ToLower(orig)
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The uppercase string is: %s\n", upper)
	fmt.Println("=======字符串的切割========")
	strsss := "The quick brown fox jumps over the lazy dog"
	sl := strings.Fields(strsss)
	fmt.Printf("Splitted in slice: %v\n", sl)
	for _, val := range sl {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str2 := "GO1|The ABC of Go|25"
	sl2 := strings.Split(str2, "|")
	fmt.Printf("Splitted in slice: %v\n", sl2)
	for _, val := range sl2 {
		fmt.Printf("%s - ", val)
	}
	fmt.Println()
	str3 := strings.Join(sl2,";")
	fmt.Printf("sl2 joined by ;: %s\n", str3)



}
