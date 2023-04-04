package main

import "fmt"

func main() {
	fmt.Println("Welcome in If Else Condition")
	loginCount := 23
	var result string
	if loginCount < 10 {
		result = "Regular user"
	} else {
		result = "Nothing"
	}

	fmt.Println("result", result)
}
