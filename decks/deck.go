package main

import ("fmt"
		"strings"
		"io/ioutil"
		"os"
		"math/rand"
		"time"
)

type deck []string

func (d deck) print () {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func newDeck() deck {
	card_values := []string{"Ace", "Two", "Three", "Four"}
	card_suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	
	cards := deck{}

	for _, suit := range card_suits {
		for _, value := range card_values {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

func deal(d deck, hand_size int) (deck, deck) {
	return d[:hand_size], d[hand_size:]
}

func (d deck) toString() string {
	
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(file_name string) {
	ioutil.WriteFile(file_name, []byte(d.toString()), 0666)
}

func newDeckFromFile(file_name string) deck {
	bs, err := ioutil.ReadFile(file_name)

	if err != nil {
		// Option1 : Log the error and return a call to newDeck()
		// Option2 : Log the error and exit the code
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source) 
	for i := range d {
		new_pos := r.Intn(len(d)-1)
		// new_pos := rand.Intn(len(d)-1)
		d[i], d[new_pos] = d[new_pos], d[i]
	}
}