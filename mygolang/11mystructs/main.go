package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	neeraj := User{"neeraj", "neeraj@poptv.sg", true, 32}

	fmt.Println(neeraj)

	fmt.Printf("Neeraj details are : %+v\n", neeraj)
	fmt.Printf("Name is %v and Email is %v ", neeraj.Name, neeraj.Email)
	// no inheritance in golang ; No super or parent

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
