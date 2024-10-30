package main

import (
	"fmt"
)

func main() {
	var N int

	
	fmt.Print("Masukkan jumlah baris segitiga bintang: ")
	_, err := fmt.Scan(&N)

	if err != nil || N <= 0 {
		fmt.Println("Input tidak valid. Harap masukkan bilangan bulat positif.")
		return
	}

	
	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println() 
	}
}