package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	cards := newDeck()
	cards.saveToFile("cards.txt")
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
	fmt.Println(cards.toString())
}
