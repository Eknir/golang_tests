package main

import (
	"io"
	"log"
	"os"
)

// create a program that reads the contents of a tet file then prints its contents to the terminal
// the file should be provided as an argument to the program when it is executed at the terminal
// to read in the arguments provided to a program you can reference the variable os.Args
// to open a file check out the doc for "os.Open"
// what interfaces does the File type implement?
// if the file type implements teh "Reader" interface, you might be able to reuse that io.Copy function!

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please provide one argument")
	}
	fileToPrint := os.Args[1]

	file, err := os.Open(fileToPrint)

	if err != nil {
		log.Fatal(err)
	}

	io.Copy(os.Stdout, file)

}
