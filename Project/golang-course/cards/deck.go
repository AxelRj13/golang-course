package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// create a new type of deck which is a slice of strings
// it's similar with OOP's classes
type deck []string

// generate new deck of cards
func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Clubs", "Hearts", "Diamonds"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	// if we declare _ on looping, then it means we don't need to use the index here
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit) // adding value to a slice
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
	fmt.Println()
}

// function return multiple values
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:] // array substring
}

// convert slice of strings into a single string value
func (d deck) toString() string {
	return strings.Join(d, ",")
}

// to export / save deck into a file
func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

// to read from file
func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ","))
}

// shuffle cards
func (d deck) shuffle() {
	// generate new seed so the random generator will always generate new numbers everytime we call shuffle (might not need in the latest GO version)
	// source := rand.NewSource(time.Now().UnixNano())
	// r := rand.New(source)

	for i := range d {
		newPosition := rand.Intn(len(d) - 1)        // would random number from 0 until n (parameter that passed to this function)
		d[i], d[newPosition] = d[newPosition], d[i] // swap position between i and newPosition
	}
}
