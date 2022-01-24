package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remaining := deal(cards, 3)
	fmt.Println("Hand")
	hand.print()
	fmt.Println("Remaining")
	remaining.print()
}

func newCard() string {
	return "Five of diamonds"
}
