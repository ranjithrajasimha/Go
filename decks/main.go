package main

func main() {
	// var card string = "Ace of Spades"
	cards := newDeckFromFile("Deck.txt")

	cards.shuffle()
	cards.print()
	// hand, remaining_deck := deal(cards, 5)

	// hand.print()
	// remaining_deck.print()

	// cards.saveToFile("Deck.txt")

	
}
