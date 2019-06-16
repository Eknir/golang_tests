package main

import "fmt"

// to whom it may concern: we have a new type called bot
// if you are a type with a function called getgreeting and you return a type string, then you are also a type of bot!
type bot interface {
	getGreeting() string
}

// hey, I have a function getgreeting and return a string! We are both bots!
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	// eb and sb are both bots, so we can pass them to printGreeting as a bot :)
	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Ola!"
}
