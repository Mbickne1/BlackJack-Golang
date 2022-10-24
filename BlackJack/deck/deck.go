package deck

import (

	"fmt"

	"math/rand"

	"time"

)

type Suit int

const (
	Hearts Suit = iota
	Diamonds
	Clubs
	Spades
)

type Card struct {
	suit	Suit 
	value	int
	name	string
}

var deck []Card 

var cardNames []string = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"} 

var values []int

func New(cardValues []int) {
	values = cardValues
	//Seed random with a different value everytime to avoid the same random number generation
	rand.Seed(time.Now().UnixNano())
	deck = generateDeck()
}

func Extend() {
	extension := generateDeck()
	deck = append(deck, extension...)
}

func generateDeck() []Card {
	//Holds local new deck
	var newDeck []Card = make([]Card, 0)
	
	for i:= 0; i < 52; i++ {
		//Initialize new card
		var newCard Card
		//The suit is equal to the floor of i / 13
		var suit Suit = Suit(i / 13)
		//Create the card
		if suit == Hearts {
			newCard = Card{Hearts, values[i % 13], cardNames[i % 13]}
		} else if suit == Diamonds {
			newCard = Card{Diamonds, values[i % 13], cardNames[i % 13]}
		} else if suit == Clubs {
			newCard = Card{Spades, values[i % 13], cardNames[i % 13]}
		} else if suit == Spades {
			newCard = Card{Clubs, values[i % 13], cardNames[i % 13]}
		}

		//Append new card to the deck
		newDeck = append(newDeck, newCard)
	}

	return newDeck
}

func Shuffle() {
	n := len(deck)
	for i, _ := range deck {
		//Generate a random number in len(deck)
		num := rand.Intn(n)
		//Swap it with current iteration card
		temp := deck[i]
		deck[i] = deck[num]
		deck[num] = temp
	}
}

func Draw() Card {
	//Draw the top card
	topCard := deck[0]
	//Slice it from the start of thea array
	deck = deck[1:]
	//Append it to the back of the array
	deck = append(deck, topCard)
	//return drawn card
	return topCard
	//return Card{Diamonds, 11 , "A"}
}

func PrintDeck() {
	var n int = len(deck);
	for i:= 0; i < n; i++ {
		var suit Suit = deck[i].suit
		var suitString string
		
		switch suit {
			case 0:	suitString = "Hearts"
			case 1: suitString = "Diamonds"
			case 2: suitString = "Spades"
			case 3: suitString = "Clubs"
		}

		fmt.Println(deck[i].name + suitString)
	}
}

func (card *Card) CardString() string {
	var suit Suit = card.suit
	var suitString string
		
	switch suit {
		case 0:	suitString = "\u0003"
		case 1: suitString = "\u0004"
		case 2: suitString = "\u0005"
		case 3: suitString = "\u0006"
	}

	return card.name + suitString
}

func (card *Card) CardValue() int {
	return card.value
}

func (card *Card) CardSuit() Suit {
	return card.suit
}

func (card *Card) CardName() string {
	return card.name
}