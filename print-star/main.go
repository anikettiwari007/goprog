package main

import (
	"math/rand"
	"strings"
	"time"
)

func main() {
	rand.NewSource(time.Now().UnixNano())

	r := rand.Intn(5) + 1

	stars := strings.Repeat("*", r)

	println(stars)
}
