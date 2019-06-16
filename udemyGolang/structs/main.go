package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firsNname string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firsNname: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email: "Jim.party@gmail.com",
			zip:   3500,
		},
	}

	// & => get the address at value (turn value into address with &)
	// give me the memory address  of the value this variable is pointing at
	// * => get the value at address (turn address into value with *)
	// give me the value this memory address is pointing at

	jim.updateName("Rinke")
	jim.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)

}

// *person is a type: pointer to person. Not an operator!
func (p *person) updateName(newFirstName string) {
	// *p is an operator (get the value at address)
	p.firsNname = newFirstName
}
