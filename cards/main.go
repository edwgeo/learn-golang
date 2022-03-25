package main

func main() {
	// var card string = "Ace of Spades"
	cards := newdeck()

	hand, remaining := deal(cards, 5)
	hand.print()
	remaining.print()
}
