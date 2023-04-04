package main

import "fmt"

func main() {
	fmt.Println("Welcome in golang")

	neeraj := User{"neeraj", "neeraj@poptv.sg", true, 32}

	fmt.Println(neeraj)

	fmt.Printf("Neeraj details are : %+v\n", neeraj)
	fmt.Printf("Name is %v and Email is %v\n ", neeraj.Name, neeraj.Email)
	neeraj.GetStatus()

	neeraj.NewEmail()

	fmt.Printf("Name is %v and Email is %v\n ", neeraj.Name, neeraj.Email)
	// no inheritance in golang ; No super or parent

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is user active: ", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test.go.dev"
	fmt.Println("Email of the user is:", u.Email)
}
