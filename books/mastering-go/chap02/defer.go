package main

import "fmt"

func d1() {
	for i := 4; i > 0; i-- {
		defer fmt.Print(i, " ")
	}
}

func d2() {

	for i := 5; i > 0; i-- {
		defer func() {
			fmt.Print(i, " ")
		}()
	}
	fmt.Println()
}

func d3() {
	for i := 6; i > 0; i-- {
		defer func(n int) {
			fmt.Print(n, " ")
		}(i)
	}
	fmt.Println()
}
func main() {
	d1()
	d2()
	d3()
}
