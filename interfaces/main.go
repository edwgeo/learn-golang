package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

// any other type in this file that matches this description (defines func getGreeting() that returns string) is also of type bot
// therefore any function that accepts type bot will also accept those types
type bot interface {
	getGreeting() string
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	// imagine that this function has some special logic that only works for englishBot
	return "Hi There!"
}

func (sb spanishBot) getGreeting() string {
	return "Hola!"
}
