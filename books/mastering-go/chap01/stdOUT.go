package main

import (
	"io"
	"os"
)

func main() {
	myString := ""
	if arguments := os.Args; len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString+"\n")
}
