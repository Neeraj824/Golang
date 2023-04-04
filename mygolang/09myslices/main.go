package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to video on Slices ")

	var fruitList = []string{"Apple", "banana", "mango"}

	fmt.Printf("Type of fruitlist is %T\n", fruitList)

	fruitList = append(fruitList, "tomato", "peach")

	fmt.Println("Fruit List is:,", fruitList)

	fruitList = append(fruitList[1:3])

	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	heighScore := make([]int, 4)

	heighScore[0] = 234
	heighScore[1] = 345
	heighScore[2] = 253
	heighScore[3] = 542
	// heighScore[4] = 777 // wrong allocation of array
	heighScore = append(heighScore, 666, 562)
	fmt.Println(heighScore)

	sort.Ints(heighScore)

	fmt.Println("heighScore", heighScore)

	// How to remove value from the slice based on index
	var courses = []string{"reastjs", "javascript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
