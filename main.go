package main

func main() {
	// var card string = "Ace of Spades"
	cards := newDeckFromFile("cards.txt")
	cards.shuffle()
	cards.print()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
