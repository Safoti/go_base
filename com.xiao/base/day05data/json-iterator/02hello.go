package main
import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"io"
	"log"
	"strings"
)
/**
 * @Author safoti
 * @Date Created in 2022/7/17
 * @Description
 **/
func main() {
	const jsonStream = `
	{"Name": "Ed", "Text": "Knock knock."}
	{"Name": "Sam", "Text": "Who's there?"}
	{"Name": "Ed", "Text": "Go fmt."}
	{"Name": "Sam", "Text": "Go fmt who?"}
	{"Name": "Ed", "Text": "Go fmt yourself!"}
`
	type Message struct {
		Name, Text string
	}
  dec:= jsoniter.NewDecoder(strings.NewReader(jsonStream))
	for  {
		var m Message
		if err := dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n", m.Name, m.Text)
	}
}
