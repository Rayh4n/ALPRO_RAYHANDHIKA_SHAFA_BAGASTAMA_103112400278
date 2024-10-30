package main

import (
	"fmt"
)

func cetakKuadrat(N int) {
	for i := 1; i <= N; i++ {
		fmt.Printf("Kuadrat dari %d adalah %d\n", i, i*i)
	}
}

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	if _, err := fmt.Scan(&N); err == nil && N > 0 {
		cetakKuadrat(N)
	} else {
		fmt.Println("Input tidak valid. Harap masukkan bilangan bulat positif.")
	}
}