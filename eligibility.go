package main

import "fmt"

func main() {
	var age int
	fmt.Print("Enter age:")
	fmt.Scan(&age)
	if age < 18 {
		fmt.Print("You are not eligible to vote")
	} else {
		fmt.Print("You are eligible to vote")
	}
}
