package main

import "fmt"

/**
 访问结构体的属性
 */


type Car struct {
	Name, Model, Color string
	WeightInKg         float64
}
//指针访问
type Employee struct {
	firstName, lastName string
	age, salary int
}


func main() {
	c := Car{Name: "Ferrari", Model: "GTC4",
		Color: "Red", WeightInKg: 1920}
	fmt.Println("Car Name: ", c.Name)
	fmt.Println("Car Color: ", c.Color)
	c.Color="Black"
	fmt.Println("Car:",c)

	emp8 :=&Employee{"Llk","gk",55,6666}
     fmt.Println("First Name:",emp8.firstName)  //两种调用方式都ok
	fmt.Println("First Name:", (*emp8).firstName)
}
