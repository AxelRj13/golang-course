package main

import "fmt"

type bot interface {
	getGreeting() string // any type that have this exact same function (including the paremeter and return type), will be a member of type bot which is predecessors/child and bot is the successor/parent
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "hello"
}

func (spanishBot) getGreeting() string {
	return "hola!"
}
