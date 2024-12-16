package main

import (
	"fmt"
	"time"
)

var (
	firstName   string    = "mohit"
	lastName    string    = "tiwari"
	currentTime time.Time = time.Now()
)

func main() {
	fmt.Println(firstName, lastName, currentTime)

	//short variable declaration
	Debug, LogLevel, startupTime := false, "info", time.Now()
	fmt.Println(Debug, LogLevel, startupTime)
}
