package main

import (
	"fmt"
	"log"

	"github.com/alEX860111/poker/card"
)

func main() {
	card, err := card.CreateCard("TS")
	fmt.Println(card.Value)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", card)
}
