package main

import (
	"fmt"
)

// Product struct dengan properti public dan private
type product struct {
	Name  string // Public (Exported)
	stock int    // Private (Unexported)
}

// new product adalah fungsi private yang mengembalikan pointer ke person
func newProduct(name string, stock int) *product {
	return &product{
		Name:  name,
		stock: stock,
	}
}

func main() {
	// Membuat instance Person menggunakan fungsi public NewPerson
	person := NewPerson("Alice", 30)

	// Mengakses properti public
	fmt.Println("Name:", person.Name)

	// Mengakses metode public
	person.Greet()

	// Mengakses properti private melalui metode public
	fmt.Println("Age:", person.Age())

	// Mengatur properti private menggunakan metode private dalam paket yang sama
	person.setAge(31)
	fmt.Println("New Age:", person.Age())

	// Membuat instance Person menggunakan fungsi public NewPerson
	product := newProduct("Alice", 30)

	// Mengakses properti public
	fmt.Println("Name:", product.Name)

	// Mencoba mengakses properti private secara langsung
	fmt.Println("Direct Age Access:", product.stock)

}
