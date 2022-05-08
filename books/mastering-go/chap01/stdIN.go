package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var f *os.File
	f = os.Stdin
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		if scanner.Text() == "exit" {
			os.Exit(0)
		}
		fmt.Println(">", scanner.Text())

	}
}
