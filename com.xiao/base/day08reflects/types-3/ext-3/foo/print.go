package foo

import (
	"fmt"
	"reflect"
	"strings"
)

/**
 * @Author safoti
 * @Date Created in 2022/7/26
 * @Description   工具类 增加指针类型
 **/
func Println(x interface{}){
	t:=reflect.TypeOf(x)
	v:=reflect.ValueOf(x)
	switch t.Kind() {
	case reflect.Struct:
		printStructExpanded(x)
	case reflect.Ptr: //指针类型
		v2 := reflect.Indirect(v)
		printStructExpanded(v2.Interface())
	default:
		fmt.Printf("Unknown type")
	}
	fmt.Print("\n")
}
/**
   如果传入的类型是结构体则进行迭代
 */
func printStructExpanded(x interface{})string  {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	if t.Kind() != reflect.Struct {
		return "<nil>"
	}

	sb := &strings.Builder{}
	n := t.NumField()
	fmt.Fprintf(sb, "%v{", t)
	for i := 0; i < n; i++ {
		tt := t.Field(i)
		vv := v.Field(i)
		fmt.Fprintf(sb, "%v: %v, ", tt.Name, vv)
	}
	fmt.Fprintln(sb, "}")

	return sb.String()

}