package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter you pizza APP")
	fmt.Println("Please rate our Pizza app 1 to 5")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	fmt.Println("Thanks for rating", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 value in rating", numRating+1)
	}
}
