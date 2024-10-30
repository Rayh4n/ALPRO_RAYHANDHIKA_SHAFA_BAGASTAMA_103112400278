package main

import "fmt"

func main() {
	var jumlahBarang int
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlahBarang)

	if jumlahBarang <= 0 {
		fmt.Println("Input tidak valid.")
		return
	}

	totalPoin := jumlahBarang*10
	if jumlahBarang > 5 {
		totalPoin += (jumlahBarang - 5) * 5
	}

	fmt.Printf("Poin yang didapatkan: %d poin\n", totalPoin)
}