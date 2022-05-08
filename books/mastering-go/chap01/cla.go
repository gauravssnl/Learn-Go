package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if arguments := os.Args; len(arguments) == 1 {
		println("Please give one or more floats.")
		os.Exit(1)
	}
	arguments := os.Args
	min, err := strconv.ParseFloat(arguments[1], 64)

	if err != nil {
		fmt.Println("Not a float:")
	}

	max, err := strconv.ParseFloat(arguments[1], 64)

	if err != nil {
		fmt.Println("Not a float")
	}

	for i := 2; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)

		if err != nil {
			fmt.Println("Not a float")
		}

		if n < min {
			min = n
		}

		if n > max {
			max = n
		}
	}

	fmt.Println("Min:", min)
	fmt.Println("Max:", max)
}
