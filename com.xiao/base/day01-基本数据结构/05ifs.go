package main

import "fmt"

func main() {
	bool1 := true
	if bool1 {
		fmt.Println("The value is true\n")
	}else {
		fmt.Printf("The value is false\n")
	}
	max:=11
	if val:=10;val>max{
		fmt.Println("大于10")
	}
		var first int = 10
		var cond int

		if first <= 0 {

			fmt.Printf("first is less than or equal to 0\n")
		} else if first > 0 && first < 5 {

			fmt.Printf("first is between 0 and 5\n")
		} else {

			fmt.Printf("first is 5 or greater\n")
		}
		if cond = 5; cond > 10 {

			fmt.Printf("cond is greater than 10\n")
		} else {

			fmt.Printf("cond is not greater than 10\n")
		}
	var num1 int = 100

	switch num1 {
	case 98, 99:
		fmt.Println("It's equal to 98")
	case 100:
		fmt.Println("It's equal to 100")
	default:
		fmt.Println("It's not equal to 98 or 100")
	}

	switch {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")
	default:
		fmt.Println("Number is 10 or greater")
	}


	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
	//for i, j := 0, N; i < j; i, j = i+1, j-1 {}
	for i:=0; i<5; i++ {
		for j:=0; j<10; j++ {
			fmt.Print(j)
		}
		fmt.Println("--")
	}
}
