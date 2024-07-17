package main

import "fmt"

func main() {
	a, b := 1, 10
	s := a + a%2
	for i := s; i <= b; i += 2 {
		fmt.Println(i)

	}

}
