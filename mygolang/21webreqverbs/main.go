package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welecome to web verb")
	// PerformGetRequest()
	// PerformPostJsonRequest()
}

func PerformGetRequest() {
	const myurl = "https://jsonplaceholder.typicode.com/posts"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length is:", response.ContentLength)

	var responseString strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	byteCount, err := responseString.Write(content)
	if err != nil {
		panic(err)
	}
	fmt.Println("byte Count is: ", byteCount)
	fmt.Println("Response : ", responseString.String())

	// fmt.Println("Print Response", string(content))

}

func PerformPostJsonRequest() {
	const myurl = "https://dummyjson.com/posts/add"

	requestBody := strings.NewReader((`
	{
		title: 'I am in love with someone.'
		userId: 69161,
	}
	
	`))

	response, err := http.Post(myurl, "applixation/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	fmt.Println("Response : ", string(content))

}

func PerformPostFormRequest() {
	const myurl = "https://dummyjson.com/posts/add"

	data := url.Values{}
	data.Add("firstname", "neeraj")
	data.Add("lastname", "srivastava")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, err := ioutil.ReadAll(response.Body)

	fmt.Println("Response : ", string(content))

}
