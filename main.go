package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/MeztliRA/gortune/quotes"
)

func main() {
	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// generate random number from a range of 0 to the length of the Quotes array
	randomNumber := rand.Intn(len(quotes.Quotes))

	// print the Quotes array using randomNumber as the index
	fmt.Println(quotes.Quotes[randomNumber])
}
