package main

import (
	"fmt"
	"log"

	"github.com/alEX860111/poker/poker"
)

func main() {
	hand, err := poker.CreateHand("2H 3C JS 2S AD")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%+v\n", hand)
}
