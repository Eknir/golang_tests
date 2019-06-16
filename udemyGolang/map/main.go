package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#745",
		"white": "#fffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for %v is %v\n", color, hex)
	}
}

//differences map and structs:
//Map:
//all keys must be of the same type, all values must be of the same type
// all keys are indexed (we can iterate over the)
// map is a reference type
// we don't need to know all keys at compile time
// use to represent a collection of related properties

//Struct:
// values can be of different type
// we can't iterate over properties of a struct
// struct is a value type
// we have to define property names + types at compile time
// use to represent a 'thing' with a lot of different properties
