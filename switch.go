package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	var op string
	fmt.Print("Enter a and b:")
	fmt.Scan(&a, &b)
	fmt.Print("Enter op: ")
	fmt.Scan(&op)
	switch op {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		fmt.Println(a / b)
	case "p":
		fmt.Println(math.Pow(a, b))
	case "m":
		fmt.Println(math.Mod(a, b))
	default:
		fmt.Println("Wrong operator")
	}
}
