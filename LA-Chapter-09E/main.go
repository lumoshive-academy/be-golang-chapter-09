package main

import (
	"be-golang-chapter-09/LA-Chapter-09E/person"
	"fmt"
)

func main() {
	// inisialisasi package
	fmt.Println("InitPerson:", person.InitPerson)
	person.SayHello()
}
