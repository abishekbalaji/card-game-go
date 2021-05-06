package main

func main() {
	// var card string = "Ace of Spades"
	cards := newDeckFromFile("cards.txt")
	cards.shuffle()
	cards.print()
}
