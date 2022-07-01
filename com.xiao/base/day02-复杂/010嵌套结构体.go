package main

import "fmt"

type Author struct {
	name,branch string
	year int
}

type HR struct {
	details Author
}

type Student struct {
	name,branch string
	year int
}

type Teacher struct {
	name,subject string
	exp int
	details Student
}


func main() {
   resul:= Author{name: "Sona",branch: "ECE",year: 2013}
	fmt.Println("\nDetails of Author")
	fmt.Println(resul)



   //调用结构体
   vt:=Teacher{
   	name: "summa",
   	subject: "java",
   	exp: 5,
   	details: Student{"Bongo", "CSE", 2},
   }
	fmt.Println("Details of the Teacher")
	fmt.Println("Teacher's name: ", vt.name)
	fmt.Println("Subject: ", vt.subject)
	fmt.Println("Experience: ", vt.exp)

	fmt.Println("\nDetails of Student")
	fmt.Println("Student's name: ", vt.details.name)
	fmt.Println("Student's branch name: ", vt.details.branch)
	fmt.Println("Year: ", vt.details.year)
}
