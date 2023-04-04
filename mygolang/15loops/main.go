package main

import "fmt"

func main() {
	fmt.Println("Welcome to loop in golangs")

	// for i := 0; i < 5; i++ {
	// 	print(i)
	// }

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// 1 Approach
	for i := 0; i < len(days); i++ {
		fmt.Println(days[i])
	}

	// 2 Approach
	for i := range days {
		fmt.Println(days[i])
	}

	// 3 Approach
	for index, day := range days {
		fmt.Printf("index is %v and value is %v\n", index, day)
	}

	// 3 Approach
	for _, day := range days {
		fmt.Printf("index is and value is %v\n", day)
	}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 1 {
			goto lco
		}
		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println(rougueValue)
		rougueValue++
	}

lco:
	fmt.Println("Jumping to goto loop")

}
