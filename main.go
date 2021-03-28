package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	quotes := [3]string{
		"I learn something every time I go into the mountains.",
		"There is nothing like a dream to create the future.",
		"The stars don't look bigger, but they do look brighter.",
	}

	number := rand.Intn(len(quotes))

	fmt.Println(quotes[number])
}
