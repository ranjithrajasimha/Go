package main

import "fmt"

func main() {
	// var card string = "Ace of Spades"
	cards := []string{"Ace of Diamonds", NewCard()}
	cards = append(cards,"Six of Spades")

	for index, card := range cards {
		fmt.Println(index, card)
	}
	
}

func NewCard() string {
	return "Ace of Hearts"
}