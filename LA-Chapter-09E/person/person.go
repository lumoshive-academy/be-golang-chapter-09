package person

import "fmt"

var InitPerson string

func init() {
	InitPerson = "my person"
}

func SayHello() {
	fmt.Println("Hello bob")
}
