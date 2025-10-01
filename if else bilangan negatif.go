package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)

	if a >= 0 {
		if a%2 == 0 {
			fmt.Println("Bukan Negatif dan Genap")
		} else {
			fmt.Println("Bukan Negatif dan Ganjil")
		}
	} else {
		if a%2 == 0 {
			fmt.Println("Genap Negatif")
		} else {
			fmt.Println("Ganjil Negatif")
		}
	}
}
