package main

import (
	"errors"
	"fmt"
	"reflect"
)

// Fungsi yang mengembalikan error
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("pembagian dengan nol")
	}
	return a / b, nil
}

func main() {
	// reflect
	var x float64 = 3.4
	fmt.Println("Tipe:", reflect.TypeOf(x))
	fmt.Println("Nilai:", reflect.ValueOf(x))

	v := reflect.ValueOf(x)
	fmt.Println("Apakah tipe float64?", v.Kind() == reflect.Float64)

	// defer
	fmt.Println("Mulai")
	defer fmt.Println("Selesai")
	fmt.Println("Tengah")

	// error handling
	result, err := divide(4, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Hasil:", result)

	// implementation recover
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Pulih dari panic:", r)
		}
	}()
	fmt.Println("Mulai")
	panic("terjadi kesalahan")
	fmt.Println("Ini tidak akan dicetak")

	// panic
	fmt.Println("Mulai")
	panic("terjadi kesalahan")
	fmt.Println("Ini tidak akan dicetak")

}
