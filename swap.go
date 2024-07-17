package main

import "fmt"

func main() {
	var a, b int
	var p, q *int
	p = &a
	q = &b
	fmt.Print("Enter a: ")
	fmt.Scan(p)
	fmt.Print("Enter b: ")
	fmt.Scan(q)
	swap(p, q)
}
func swap(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
	fmt.Printf("a = %d, b = %d", *a, *b)
}
