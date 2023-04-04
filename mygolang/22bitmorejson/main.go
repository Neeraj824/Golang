package main

import (
	"encoding/json"
	"fmt"
)

type courses struct {
	Name     string   `json:"coursename"`
	Price    int      `json:"price"`
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Wecome to json video in Golang")

	// EncodeJson()
	DecodeJson()

}

func EncodeJson() {
	nCourses := []courses{
		{"ReactJS Bootcamp", 299, "neeraj.com", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "neeraj.com", "bcc123", []string{"full-stack", "js"}},
		{"ANGULAR Bootcamp", 299, "neeraj.com", "sbc123", nil},
	}

	// package this data as json code

	// finalJson, err := json.Marshal(nCourses)
	finalJson, err := json.MarshalIndent(nCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "MERN Bootcamp",
		"price": 199,
		"website": "neeraj.com",
		"tags": ["full-stack","js"]
	}
	 `)

	// Ist Approach
	var lcoCourse courses

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)
	}

	// IIst Approach
	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("key is %v and Value is %v and Type is: %T\n", k, v, v)
	}

}
