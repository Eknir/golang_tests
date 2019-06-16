package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreetingEB(eb)
	printGreetingSB(sb)
}

// really want to re-use print greeting for both bots (same logic)
func printGreetingEB(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func printGreetingSB(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func (englishBot) getGreeting() string {
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Ola!"
}
