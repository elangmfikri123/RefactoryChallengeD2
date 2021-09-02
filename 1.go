package main

import "fmt"

//odd and even numbers
func main() {
	for i := 1; i < 10; i++ {
		fmt.Print("Masukkan Bilangan : ")
		fmt.Scan(&i)
		if i <= 10 {
			if i%2 == 0 {
				fmt.Print(i, " Bilangan Genap \n")
			} else {
				fmt.Print(i, " Bilangan Ganjil \n")
			}
		} else {
			fmt.Print(" Selesai, Angka lebih dari 10 \n")
		}
	}

}
