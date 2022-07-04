package main

import "fmt"

/**
  结构体番外篇
 */
func main() {
	//创建匿名结构体

	emp3 :=struct {
		firstName string
		lastName string
		age int
		salary int
	}{
		firstName: "kj",
		lastName: "kr",
		age: 15,
		salary: 200,
	}
	fmt.Println("Employee 3", emp3)
	//访问对象的属性
	fmt.Println("firstName 3", emp3.firstName)
	fmt.Println("lastName 3", emp3.lastName)
	fmt.Println("age 3", emp3.age)
	fmt.Println("salary 3", emp3.salary)



	//没有显示的设置初始化值
  	var emp4  Employee
	//获取默认属性
	fmt.Println("First Name:", emp4.firstName)
	fmt.Println("Last Name:", emp4.lastName)
	fmt.Println("Age:", emp4.age)
	fmt.Println("Salary:", emp4.salary)


	//赋予部分值
	emp5 := Employee{
		firstName: "John",
		lastName:  "Paul",
	}
	fmt.Println("First Name:", emp5.firstName)
	fmt.Println("Last Name:", emp5.lastName)
	fmt.Println("Age:", emp5.age)
	fmt.Println("Salary:", emp5.salary)


	e := Employee{
		firstName: "Mark Andrew",
		age:  50,
	}

	fmt.Printf("\n\nEmployee age before change: %d", e.age)
	e.changeAge(54)
	fmt.Printf("\nEmployee age after change: %d", e.age)


}

func (e *Employee)changeAge(newAge int){
	e.age = newAge
}

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}