package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This need to fo in a file - LearnCodeOnline.in"

	file, err := os.Create("./neerajfile.txt")

	// if err != nil {
	// 	panic(err)
	// }
	checkNilErr(err)
	length, err := io.WriteString(file, content)

	checkNilErr(err)
	fmt.Println("Length is :", length)

	defer file.Close()

	readFile("./neerajfile.txt")
}

func readFile(fileName string) {
	databyte, err := ioutil.ReadFile(fileName)

	checkNilErr(err)
	fmt.Println("Text Data inside the file: \n", string(databyte))
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
