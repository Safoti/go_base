package main

import "fmt"

/**
 * @Author safoti
 * @Date Created in 2022/7/24
 * @Description 
 **/
func main() {
	//mystruct.number=3.14
	da:=Date{}
	da.setYead(2019)
	fmt.Println(da.Year)
}

var mystruct struct{
	number float64
	world string
	toggle bool
}

type Date struct {
	Year int
	Month int
	Day int
}

func (d *Date)setYead(year int) {
	d.Year=year
}