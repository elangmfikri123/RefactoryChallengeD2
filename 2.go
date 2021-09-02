package main

import "fmt"

//grade from score
func main() {
	fmt.Print("Masukkan Nilai: ")
	var score float32
	fmt.Scanln(&score)

	if score > 3.10 && score <= 4 {
		fmt.Println("Nilai A")
	} else if score > 2.10 && score <= 3 {
		fmt.Println("Nilai B")
	} else if score > 1.60 && score <= 2 {
		fmt.Println("Nilai C")
	} else if score > 0 && score <= 1.50 {
		fmt.Println("Nilai D")
	} else if score == 0 {
		fmt.Println("Nilai E")
	} else {
		fmt.Println("Masukkan nilai yang tepat")
	}
}
