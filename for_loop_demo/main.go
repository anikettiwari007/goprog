package main

import "fmt"

func main() {

	//simple loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//looping over arrays and slices
	names := []string{"Jim", "Jane", "Joe", "June"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	//looping over a map
	config := map[string]string{
		"debug":    "1",
		"logLevel": "warn",
		"version":  "1.2.1",
	}

	for key, value := range config {
		fmt.Println(key, "=", value)
	}

	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}
	var index string
	max := 0
	for key, value := range words {
		if value > max {
			max = value
			index = key
		}
	}
	fmt.Println("Most popular word: ", index)
	fmt.Println("With a count of: ", words[index])
}
