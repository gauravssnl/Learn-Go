package main

import "fmt"

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false)
	res := func() bool {
		return true || true
	}
	fmt.Println(res, res())
	fmt.Println(!true)
}
