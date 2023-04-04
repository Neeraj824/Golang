package main

import "fmt"

func main() {
	// defer work LILO(Last IN Last OUT)
	defer fmt.Println("World")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}

// Output
// Hello
// 43210Two
// One
// World
