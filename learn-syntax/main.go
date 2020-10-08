package main

import "fmt"

func main() {
	// hello()
	// fmt.Println(fmt.Sprintf("Flag: ", ifStatement(15,10)))
	fmt.Println("Flag:", ifStatement(25,20))
}

func hello() {
	fmt.Println("Hai")
}

func ifStatement (x int, y int) bool {
	if ok := x < y;ok {
		return ok
	} else {
		return ok
	}
}

