package main

import (
	"fmt"
	"strconv"
)

func getOptions() (int, float64) {
	var decksInput string
	var decks int
	isGettingNumDecks := true

	for isGettingNumDecks == true {
		fmt.Print("How many decks should the game use (1-6)? ")
		fmt.Scanln(&decksInput)

		var err error
		decks, err = strconv.Atoi(decksInput)
		if err != nil || decks < 1 || decks > 6 {
			fmt.Println("Invalid input, please input a number.")
			continue
		}

		isGettingNumDecks = false
	}

	var deckPenetrationInput string
	var deckPenetration int
	isGettingDeckPenetration := true

	for isGettingDeckPenetration == true {
		fmt.Print("What should the deck penetration percentage be (50-80)? ")
		fmt.Scanln(&deckPenetrationInput)

		var err error
		deckPenetration, err = strconv.Atoi(deckPenetrationInput)
		if err != nil || deckPenetration < 50 || deckPenetration > 80 {
			fmt.Println("Invalid input, please input a number (50-80).")
			continue
		}

		isGettingDeckPenetration = false
	}

	return decks, float64(deckPenetration) / 100
}

func main() {
	decks, deckPenetration := getOptions()
	deck := NewMultipleDecks(decks, deckPenetration)
	stats := PlayGame(deck)
	if stats.Wins > 0 || stats.Losses > 0 {
		fmt.Printf("\nGame Statistics: %+v\n", stats)
	}
}
