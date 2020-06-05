package main

import "fmt"

func (hand *Hand) deal(deck Deck, i *uint8) {
	newCard := deck[*i]
	if newCard.Type == "Ace" {
		hand.HasAce = true
	}
	hand.Cards = append(hand.Cards, newCard)
	hand.Total = hand.Total + newCard.Value
	*i++
}

// Stats represents the statistics for a shoe
type Stats struct {
	Wins    int
	Losses  int
	WinRate float64
}

// Hand consists of the cards in the hand and the total value of those cards
type Hand struct {
	Cards    []Card
	HasAce   bool
	IsDealer bool
	Total    uint8
}

func (hand *Hand) displayHand(showAll bool) {
	var formattedHandSlice []string
	cards := hand.Cards

	if hand.IsDealer && len(cards) == 2 && !showAll {
		// first card is hidden for dealer
		formattedHandSlice = append(formattedHandSlice, cards[1].Type)
	} else {
		for _, card := range cards {
			formattedHandSlice = append(formattedHandSlice, card.Type)
		}
	}

	if hand.IsDealer {
		fmt.Printf("Dealer Hand: %v\n", formattedHandSlice)
	} else {
		fmt.Printf("Player Hand: %v\n", formattedHandSlice)
	}
}

func (hand *Hand) getHighestTotal() uint8 {
	if hand.Total > 21 {
		// no reason to adjust, it's a bust anyway
		return hand.Total
	}

	if hand.HasAce {
		// use 11 value of an ace instead of 1
		possibleTotal := hand.Total + 10
		if possibleTotal <= 21 {
			return possibleTotal
		}
	}

	return hand.Total
}

func (hand *Hand) hasBlackjack() bool {
	return (len(hand.Cards) == 2 && hand.HasAce && hand.getHighestTotal() == 21)
}

// PlayGame plays a game given a set Deck
func PlayGame(deck Deck) Stats {
	var (
		i      uint8 = 1 // burn the first card
		wins   int
		losses int
	)

	fmt.Println("Welcome to the table!")

	// I need a flag to stop after the cut card, and I need to not have it dealt
	for {
		var dealerHand Hand
		dealerHand.IsDealer = true
		var playerHand Hand

		playerHand.deal(deck, &i)
		dealerHand.deal(deck, &i)
		playerHand.deal(deck, &i)
		dealerHand.deal(deck, &i)

		dealerHand.displayHand(false)
		playerHand.displayHand(true)

		dealerHasBlackjack := dealerHand.hasBlackjack()
		playerHasBlackjack := playerHand.hasBlackjack()
		takeEvenMoney := true

		var takeEvenMoneyChoice string
		if dealerHand.Cards[1].Type == "ACE" && playerHasBlackjack {
			for {
				fmt.Println("Would you like even money? [y: yes, n: no]")
				fmt.Scanln(&takeEvenMoneyChoice)
				if takeEvenMoneyChoice == "y" {
					takeEvenMoney = true
					break
				} else if takeEvenMoneyChoice != "n" {
					fmt.Println("Invalid input.")
				}
			}
		}

		if dealerHasBlackjack && playerHasBlackjack {
			if takeEvenMoney {
				fmt.Println("Paying even money.")
				wins = wins + 1
				continue
			}
			fmt.Println("Push.")
			continue
		}

		if dealerHasBlackjack {
			fmt.Println("Dealer wins with blackjack.")
			losses = losses + 1
			continue
		}

		if playerHasBlackjack {
			fmt.Println("Player wins with blackjack.")
			wins = wins + 1
			continue
		}

		var choice string
		playerHighestTotal := playerHand.getHighestTotal()
		for playerHighestTotal < 21 {
			fmt.Print("Pick an option: [h: hit, s: stand, d: double, sp: split, e: end game]: ")
			fmt.Scanln(&choice)

			if choice == "e" {
				fmt.Println("Thank you for playing.")
				return Stats{
					Wins:    wins,
					Losses:  losses,
					WinRate: float64(wins) / (float64(wins) + float64(losses)),
				}
			}

			if choice == "s" {
				break
			} else if choice == "h" {
				playerHand.deal(deck, &i)
				playerHighestTotal = playerHand.getHighestTotal()
				dealerHand.displayHand(false)
				playerHand.displayHand(true)
			} else if choice == "d" {
				// Add Double Logic
			} else if choice == "sp" {
				// Add Split Logic
			} else {
				fmt.Println("Invalid input.")
				continue
			}
		}

		dealerHighestTotal := dealerHand.getHighestTotal()
		for dealerHighestTotal < 21 {
			// assume dealer hits soft 17
			if dealerHighestTotal < 17 || (dealerHighestTotal == 17 && dealerHand.HasAce) {
				dealerHand.deal(deck, &i)
				dealerHighestTotal = dealerHand.getHighestTotal()
				continue
			}
			break
		}

		playerTotal := playerHand.getHighestTotal()
		dealerTotal := dealerHand.getHighestTotal()
		dealerHand.displayHand(true)
		playerHand.displayHand(true)
		if playerTotal > 21 || (dealerTotal <= 21 && (dealerTotal > playerTotal)) {
			fmt.Println("Dealer wins.")
			losses = losses + 1
		} else if dealerTotal > 21 || playerTotal > dealerTotal {
			fmt.Println("Player Wins.")
			wins = wins + 1
		} else {
			fmt.Println("Push.")
		}
	}
}
