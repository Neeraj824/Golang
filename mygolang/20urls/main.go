package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://loc.dev:3000/learn?coursename=reactjs&payment_id=234234234"

func main() {
	fmt.Println("Welcome to handle URL's in golang")
	fmt.Println("myurl", myurl)

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparam := result.Query()

	fmt.Printf("The type of query params are: %T\n", qparam)

	fmt.Println(qparam["coursename"])
	fmt.Println(qparam["payment_id"])

	for _, val := range qparam {
		fmt.Println("Param is: ", val)
	}

	partsOfUrls := &url.URL{
		Scheme:   "https",
		Host:     "loc.dev",
		Path:     "/learn",
		RawQuery: "user=neeraj",
	}

	anotherURl := partsOfUrls.String()
	fmt.Println("Another URL", anotherURl)
}
