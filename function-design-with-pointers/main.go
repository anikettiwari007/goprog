package main

import "fmt"

func add5Value(count int) {
	count += 5
	fmt.Println("addValue: ", count)
}

func addPointer(count *int) {
	*count += 5
	fmt.Println("addPointer: ", *count)
}
func main() {
	var count int
	add5Value(count)
	addPointer(&count)
	fmt.Println("add5Point post: ", count)
}
