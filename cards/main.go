package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	cards := []string{newCard(), newCard()}

	for i, card := range cards {
		fmt.Println(i, card)
	}

}

func newCard() string {
	return "Five of Diamonds"
}
