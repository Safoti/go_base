package main

import (
	"fmt"
	"reflect"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/28
 * @Description 
 **/
type (
	// TableField describe a RDBMS table field's value
	TableField struct {
		Name  string // ex: Lname, Fname, Age, etc.
		Type  string // ex: int, string, float64, etc
		Value string // all values stored as string
	}
)
var (
	row1 = []TableField{
		{Name: "ID", Type: "int", Value: "1"},
		{Name: "Fname", Type: "string", Value: "Jane"},
		{Name: "Lname", Type: "string", Value: "Doe"},
		{Name: "Age", Type: "uint8", Value: "41"},
		{Name: "Height", Type: "float64", Value: "5.5"},
	}
	row2 = []TableField{
		{Name: "ID", Type: "int", Value: "2"},
		{Name: "Fname", Type: "string", Value: "Mark"},
		{Name: "Lname", Type: "string", Value: "Smith"},
		{Name: "Age", Type: "uint8", Value: "23"},
		{Name: "Height", Type: "float64", Value: "5.9"},
	}
	row3 = []TableField{
		{Name: "ID", Type: "int", Value: "3"},
		{Name: "Fname", Type: "string", Value: "Anne"},
		{Name: "Lname", Type: "string", Value: "Marie"},
		{Name: "Age", Type: "uint8", Value: "89"},
		{Name: "Height", Type: "float64", Value: "5.2"},
	}
	// PeopleData is our collection of data populated from RDBMS
	PeopleData = [][]TableField{row1, row2, row3}
)

// dumpData print out our data to see what we have
func dumpData(d [][]TableField) {
	var spacing = []int{7, 16, 16, 13, 18}
	headers := d[0]
	var col string
	for i, c := range headers {
		col = fmt.Sprintf("%v(%v) |", c.Name, c.Type)
		fmt.Printf("%*v", spacing[i], col)
	}
	fmt.Println()

	for _, r := range d {
		for i, c := range r {
			fmt.Printf("%*v |", spacing[i], c.Value)
		}
		fmt.Println()
	}
}

func main() {
	dumpData(PeopleData)

	s1 := dataToStruct(PeopleData[0])
	fmt.Printf("%#v\n", s1)
}
func dataToStruct(data []TableField) interface{} {
	var structFields []reflect.StructField
	structFields =make([]reflect.StructField,len(data))
	for i, tf := range data {
		structFields[i].Name=tf.Name
	}
	typ:=reflect.StructOf(structFields)
	v:=reflect.New(typ).Elem()
	s:=v.Addr().Interface()
	return s
}