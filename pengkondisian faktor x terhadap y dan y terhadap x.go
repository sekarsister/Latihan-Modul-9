package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	if a == b {
		if a%b == 0 {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}

		if b%a == 0 {
			fmt.Println("true")
		} else {
			fmt.Println("false")
		}
	} else {
		if a%b == 0 {
			fmt.Println("false")
		} else {
			fmt.Println("true")
		}

		if b%a == 0 {
			fmt.Println("false")
		} else {
			fmt.Println("true")
		}
	}

}
