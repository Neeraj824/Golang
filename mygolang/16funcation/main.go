package main

import "fmt"

func main() {
	fmt.Println("Welcome to function in golang")
	greeter()
	result := adder(3, 5)
	fmt.Println("Result is: ", result)
	greeterTwo()

	resultU, mymessage := proAdder(3, 5, 2, 5, 4)
	fmt.Println("Result is ", resultU)
	fmt.Println("Pro Message is ", mymessage)
}

func greeter() {
	fmt.Println("Nasmtey from golang")
}
func greeterTwo() {
	fmt.Println("Another method")
}

func adder(valOne int, valTwo int) int {
	return (valOne + valTwo)
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Hi Pro Result function"
}
