package main

import "fmt"

func main() {
	// Contoh penggunaan general placeholder
	str := "hello"
	num := 42
	fmt.Printf("String: %s, Number: %d\n", str, num)

	// Integer formatting
	fmt.Printf("Binary: %b, Decimal: %d, Octal: %o, Hex: %x\n", num, num, num, num)
	// Floating point formatting
	floatNum := 3.14159
	fmt.Printf("Scientific: %e, Float: %f, Shorter: %g\n", floatNum, floatNum, floatNum)

	// Struct formatting
	person := struct {
		Name string
		Age  int
	}{"Alice", 30}
	fmt.Printf("Struct: %v\n", person)
	fmt.Printf("Struct with field names: %+v\n", person)
	fmt.Printf("Go-syntax struct: %#v\n", person)

	// Type formatting
	fmt.Printf("Type of num: %T\n", num)
	fmt.Printf("Type of person: %T\n", person)

	// Pointer formatting
	fmt.Printf("Pointer: %p\n", &num)
}
