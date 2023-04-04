package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")
	presentTime := time.Now()

	fmt.Println("PresentTime", presentTime.Format("2006-01-02T15:04:05"))
}
