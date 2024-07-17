package main

import "fmt"

func main() {
	var x string
	fmt.Print("Enter a or b:")
	fmt.Scan(&x)
	if x == "a" {
		fmt.Println(biodata())
	}
	if x == "b" {
		fmt.Println(hobbies())
	}

}
func biodata() string {
	return "My name is Susmitha Reddy currently pursuing btech II year in Shri Vishnu Engineering College for women.I am interested in learning new technologies."
}
func hobbies() string {
	return "My hobbies are dancing,watching movies"
}
