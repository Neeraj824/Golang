package main

import "fmt"

func main() {
	fmt.Println("Maps in golangs")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["PY"] = "Python"
	languages["RB"] = "RUBY"

	fmt.Println(languages)

	delete(languages, "RB")
	fmt.Println(languages)

	for key, value := range languages {
		fmt.Printf("For key %v , value is %v\n", key, value)
	}

}
