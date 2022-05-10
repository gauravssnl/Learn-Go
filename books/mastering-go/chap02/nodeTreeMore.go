package main

import "fmt"

func main() {
	varOne := 1
	varTwo := 2
	fmt.Println("Hello there!")
	functionOne(varOne)
	functionOne((varTwo))

}

func functionOne(x int) {
	fmt.Println(x)
}
