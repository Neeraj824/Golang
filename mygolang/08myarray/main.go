package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golangs")

	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Peach"

	fmt.Println("Fruit List is: ", fruitList)
	fmt.Println("Length Fruit List is: ", len(fruitList))

	var vegList = [3]string{"potato", "beans", "mushroom"}
	fmt.Println("Vegy list is: ", vegList)
	fmt.Println("Length of Vegy list is: ", len(vegList))

}
