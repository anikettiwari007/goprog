package main

import (
	"fmt"
	"time"
)

func main() {
	Debug, LogLevel, startupTime := getConfig()

	fmt.Println(Debug, LogLevel, startupTime)
}

func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}
