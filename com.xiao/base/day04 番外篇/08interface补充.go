package main

import "fmt"

type humanInterface interface {
	eat() string
	play() string
}
type man struct {
	name string
}
func (p man )eat()string  {
	return "eat banana"
}
func(p man)play()string{
	return "play game"
}



/**
  只实现其中一个方法

 */

type dogInterface interface {
	eat() string
	play() string
}

type dog1 struct {
	name string
}

func (d dog1) eat() string {
	return "Eat dog food"
}

func main() {
var human humanInterface
 human=new(man)
 fmt.Println(human.eat())
 fmt.Println(human.play())

 
 //var dogs dogInterface
 //dogs=new(dog1)

}
