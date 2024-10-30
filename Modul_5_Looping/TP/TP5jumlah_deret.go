package main

import "fmt"

func main()  {
	var n int = 3
	fmt.Print("Masukkan nilai n; ")
	fmt.Scan(&n)

	jumlah := 0
	for i := 1; i <= n; i++ {
        jumlah += i
	
	}

    fmt.Printf("Jumlah deret angka dari 1 hingga %d adalah %d\n", n, jumlah)
}
