package main

import (
	"fmt"
)

func main() {
	var bilangan int

	fmt.Print("Masukkan bilangan bulat: ")
	fmt.Scan(&bilangan)

	isGenapNegatif := (bilangan < 0) && (bilangan%2 == 0)

	fmt.Println(isGenapNegatif)
}