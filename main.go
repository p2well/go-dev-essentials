package main

import (
	"fmt"

	"github.com/p2well/go-dev-essentials/mascot"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Welcome Gophers 😎")
	fmt.Println(mascot.BestMascot())
	fmt.Println(quote.Go())
}
