package main

import (
	"math"
	"math/rand"
	"time"
)

// Card holds the card suits and types in the deck
type Card struct {
	Type  string
	Suit  string
	Value uint8
}

// Deck consists of all the Cards in the deck
type Deck []Card

var cardTypes = []string{
	"Ace", "Two", "Three", "Four", "Five", "Six", "Seven",
	"Eight", "Nine", "Ten", "Jack", "Queen", "King",
}
var suits = []string{"Heart", "Diamond", "Club", "Spade"}

func getCardValue(cardType string) uint8 {
	switch cardType {
	case "Two":
		return 2
	case "Ace":
		return 1
	case "Three":
		return 3
	case "Four":
		return 4
	case "Five":
		return 5
	case "Six":
		return 6
	case "Seven":
		return 7
	case "Eight":
		return 8
	case "Nine":
		return 9
	default:
		return 10
	}
}

// NewDeck generates a new default deck of cards
func NewDeck() (deck Deck) {
	for i := 0; i < len(cardTypes); i++ {
		for j := 0; j < len(suits); j++ {
			cardType := cardTypes[i]
			card := Card{
				Type:  cardType,
				Suit:  suits[j],
				Value: getCardValue(cardType),
			}
			deck = append(deck, card)
		}
	}

	deck.shuffle()
	return deck
}

var cutCard = Card{
	Type:  "CUT_TYPE",
	Suit:  "CUT_SUIT",
	Value: 0,
}

func (deck Deck) insertCutCard(cutLocation uint8) Deck {
	deck = deck[0 : len(deck)+1]
	copy(deck[cutLocation+1:], deck[cutLocation:])
	deck[cutLocation] = cutCard
	return deck
}

// NewMultipleDecks generates n decks as one output Deck
func NewMultipleDecks(n int, deckPenetration float64) (deck Deck) {
	for i := 0; i < n; i++ {
		deck = append(deck, NewDeck()...)
	}
	deck.shuffle()
	cutLocation := uint8(math.Floor(float64(len(deck)) * deckPenetration))
	newDeck := deck.insertCutCard(cutLocation)
	return newDeck
}

func (deck Deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	for i := range deck {
		j := rand.Intn(i + 1)
		deck[i], deck[j] = deck[j], deck[i]
	}
}
