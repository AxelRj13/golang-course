package main

import "fmt"

type bot interface {
	getGreeting() string // any type that have this exact same function (including the paremeter and return type), will be a member of type bot which is predecessors/child and bot is the successor/parent
}

type englishBot struct {
	greeting string
}

type spanishBot struct {
	greeting string
}

func main() {
	eb := englishBot{greeting: "Hello"}
	sb := spanishBot{greeting: "Hola!"}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// these functions are like overidding from the bot interface with the custom function each
func (eb englishBot) getGreeting() string {
	return eb.greeting
}

func (sb spanishBot) getGreeting() string {
	return sb.greeting
}
