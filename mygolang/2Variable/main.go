package main

import (
	"fmt"
)

const LoginToken string = "sdfsdfsd" // Public

func main() {
	var username string = "Neeraj"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T\n", username)
	// fmt.Println(reflect.TypeOf(username))

	var firstname = "neeraj"
	fmt.Println(firstname)

	lastname := "srivastava"
	fmt.Println(lastname)

	fmt.Println("LoginToken", LoginToken)

}
