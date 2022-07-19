package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//s:="hello.name"
   //strings.Split(s,".")
	//测试下方法的调用
	//lz:=lktbz{"lk",15}
	//name := lz.getName(15)
	//fmt.Println(name)
	//Trims()
	//ToUPAndLow()
	//Split()
	//Replaceat()
    //IndexDemo()
	 CountDemo()
}
func CountDemo(){
	fmt.Println(strings.Count("cheese", "e"))
	fmt.Println(strings.Count("five", "")) // before & after each rune


	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))


	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("fail", "ui"))
	fmt.Println(strings.ContainsAny("ure", "ui"))
	fmt.Println(strings.ContainsAny("failure", "ui"))
	fmt.Println(strings.ContainsAny("foo", ""))
	fmt.Println(strings.ContainsAny("", ""))



}

func IndexDemo() {
	fmt.Println(strings.Index("chicken", "ken"))
	fmt.Println(strings.Index("chicken", "dmr"))

	fmt.Println(strings.IndexAny("chicken", "aeiouy"))
	fmt.Println(strings.IndexAny("crwth", "aeiouy"))

	fmt.Println(strings.IndexByte("golang", 'g'))
	fmt.Println(strings.IndexByte("gophers", 'h'))
	fmt.Println(strings.IndexByte("golang", 'x'))

	fmt.Println(strings.IndexRune("chicken", 'k'))
	fmt.Println(strings.IndexRune("chicken", 'd'))

	s := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(s, ", "))

	fmt.Println(strings.Index("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "go"))
	fmt.Println(strings.LastIndex("go gopher", "rodent"))


	show := func(s, sep string) {
		before, after, found := strings.Cut(s, sep)
		fmt.Printf("Cut(%q, %q) = %q, %q, %v\n", s, sep, before, after, found)
	}
	show("Gopher", "Go")
	show("Gopher", "ph")
	show("Gopher", "er")
	show("Gopher", "Badger")
}
func Replaceat()  {
	fmt.Println("ba" + strings.Repeat("na", 2))
	fmt.Println(strings.Repeat("jji", 2))

	fmt.Println(strings.Replace("oink oink oink", "k", "ky", 2))
	fmt.Println(strings.Replace("oink oink oink", "oink", "moo", -1))


	fmt.Println(strings.ReplaceAll("oink oink oink", "oink", "moo"))

}
func Split(){
	fmt.Printf("%q\n", strings.Split("a,b,c", ","))

	fmt.Printf("%q\n", strings.Split("a man a plan a canal panama", "a "))

	fmt.Printf("%q\n", strings.Split(" xyz ", ""))
	fmt.Printf("%q\n", strings.Split("", "Bernardo O'Higgins"))


	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2))


	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2))
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil)

}
func ToUPAndLow()  {
	fmt.Println(strings.ToLower("Gopher"))

	fmt.Println(strings.ToTitle("her royal highness"))
	fmt.Println(strings.ToTitle("loud noises"))
	fmt.Println(strings.ToTitle("хлеб"))

	fmt.Println(strings.ToUpper("Gopher"))

}
/**
   Trim 相关
 */
func Trims(){
	fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!","!¡"))


	fmt.Println(strings.TrimLeft("¡¡¡Hello, Gophers!!!","!¡"))

	fmt.Println(strings.TrimLeftFunc("¡¡¡Hello, Gopherss!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
	fmt.Println(strings.TrimLeftFunc("¡¡¡222!!!", func(r rune) bool {
		return unicode.IsLetter(r) && unicode.IsNumber(r)
	}))



	var s = "¡¡¡Hello, Gophers!!!"
	s=strings.TrimPrefix(s,"¡¡¡Hello,") //去掉匹配的内容
	fmt.Println(s)
	var ssd = "¡¡¡Hello, Gophers!!!"
	ssd=strings.TrimPrefix(ssd,"¡¡¡Howdy, ")
	fmt.Println(ssd)
	fmt.Println(strings.TrimRight("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println(strings.TrimRight("¡¡¡Hello, Gophers!!!", "s!!!"))



	//去除空格
	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))
	fmt.Println(strings.TrimSpace("  Hello,      Gophers    "))
	//去除后缀
	var ddss = "¡¡¡Hello, Gophers!!!"
	ddss = strings.TrimSuffix(ddss, ", Gophers!!!")
    fmt.Println(ddss)

}
func MyFuncDemo(fi,se string)string{
	if fi=="" ||se ==""{
		return fi
	}
	return ""
}
//传递一个函数进来当参数
func TrimLeftFunc(s string, f func(rune) bool) string {

	return ""
}


//定义一个方法
func (u *lktbz)getName(age int)string{
	//check u
    //chekc age

	if age ==u.Age{
		fmt.Println(u.Name)
	   return u.Name
	}
	return ""
}
type lktbz struct {
	 Name string
	 Age int
}