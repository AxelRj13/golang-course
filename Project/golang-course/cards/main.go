package main

func main() {
	// card := newCard() // instantiate a variable using function
	// fmt.Println(card)
	// printState()            // calling function from other files in the same package
	// card = "Three of Clubs" // assign new value to the variable that already instantiated
	// fmt.Println(card)
	// fmt.Println()

	// slices and loops
	// cards := deck{newCard(), "King of Hearts", "Ace of Spades"} // assign value to slice of strings that equals to the 'deck' type
	// cards = append(cards, "Six of Spades") // add new value to the slices

	// assign using newDeck function
	// cards := newDeck()

	// New deck from file
	cards := newDeckFromFile("MyCards.txt")
	cards.shuffle()

	// call function from deck type (class) that doesn't return anything (void)
	cards.print()

	// This function deal will return multiple values and this is how we capture it
	// hand, remainingDeck := deal(cards, 5)
	// fmt.Println("Cards on hand")
	// hand.print()
	// fmt.Println("Remaining cards on deck")
	// remainingDeck.print()

	// string to byte conversion
	// cards.saveToFile("MyCards.txt")

	// test example
	// call function from laptopSize type that return a value of data type
	// fmt.Println()
	// size := laptopSize(1.5)                            // instantiate a value of laptopSize type
	// fmt.Println("LaptopSize:", size.getSizeOfLaptop()) // print out the return value from laptopSize function
}

func newCard() string {
	return "Five of Diamonds"
}
