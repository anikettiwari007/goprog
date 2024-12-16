package main

import "fmt"

func swap(a *int, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 5
	b := 6
	fmt.Println("------before----")
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	swap(&a, &b)
	fmt.Println("------after----")
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
}
