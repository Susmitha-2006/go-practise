package main

import "fmt"

func main() {
	var a, c int

	fmt.Print("Enter a: ")
	fmt.Scan(&a)
	for i := 2; i < a/2; i++ {
		if a%i == 0 {
			c = 1
			break
		}
	}
	if c == 1 {
		fmt.Println("It is not prime")
	} else {
		fmt.Println("It is prime")
	}
}
