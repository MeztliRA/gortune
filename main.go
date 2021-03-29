package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/MeztliRA/gortune/quotes"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	randomNumber := rand.Intn(len(quotes.Quotes))

	fmt.Println(quotes.Quotes[randomNumber])
}
