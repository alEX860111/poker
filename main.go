package main

import (
	"fmt"
	"log"

	"./card"
)

func main() {
	card, err := card.CreateCard("TS")
	fmt.Println(card.Value)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", card)
}
