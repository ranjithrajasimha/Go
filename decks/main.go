package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	card := NewCard()
	fmt.Println(card)
}

func NewCard() string {
	return "Ace of Spades"
}