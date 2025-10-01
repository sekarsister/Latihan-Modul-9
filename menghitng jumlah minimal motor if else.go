package main

import "fmt"

func main() {
	var jumlahOrang int
	fmt.Scan(&jumlahOrang)

	jumlahMotor := 0
	sisaOrang := jumlahOrang

	for sisaOrang > 0 {
		if sisaOrang >= 2 {
			sisaOrang -= 2
		} else {
			sisaOrang -= 1
		}
		jumlahMotor++
	}

	fmt.Println(jumlahMotor)
}
