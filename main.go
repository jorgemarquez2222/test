package main

import (
	"fmt"

	"github.com/jorgemarquez2222/test/greet"
	"rsc.io/quote"
)

func main() {
	fmt.Print(greet.English())
	fmt.Println(quote.Hello())
}
