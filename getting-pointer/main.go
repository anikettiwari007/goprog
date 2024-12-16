package main

import (
	"fmt"
	"time"
)

func main() {
	//declare a pointer variable using var
	var count1 *int
	// declare a pointer variable using new keyword
	count2 := new(int)

	//create a pointer from the existing variable
	countTemp := 5
	count3 := &countTemp

	//create a pointer from some types without a temporary variable
	t := &time.Time{}

	fmt.Printf("count1: %#v\n", count1)
	fmt.Printf("count2: %#v\n", count2)
	fmt.Printf("count3: %#v\n", count3)
	fmt.Printf("time : %#v\n", t)
}
