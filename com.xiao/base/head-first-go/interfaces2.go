package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/25
 * @Description 
 **/

type NoiseMaker interface {
	MakeSound()

}

type Whistle string


func (w Whistle)MakeSound()  {
	fmt.Println("Whistle !")
}

type Horno string

func (w Horno)MakeSound()  {
	fmt.Println("Horno !")
}
type Robot string

func (w Robot)MakeSound()  {
	fmt.Println("Robot !")
}

func (w Robot)Walking()  {
	fmt.Println("Walking !")
}
func main() {
	//var toy NoiseMaker
	//toy=Whistle("toyo")
	//toy.MakeSound()
	//
	//toy=Horno("ho")
	//toy.MakeSound()


	ppls(Whistle("toyo"))	
	ppls(Horno("ho"))

}
func ppls(  niose NoiseMaker){
	niose.MakeSound()
	//niose.Walking() 接口中并没有此方法
}