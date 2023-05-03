package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?username=himanshu&oderid=rksgdnvs"

func main() {
	fmt.Println("Urls in golang")

	//parsing
	result, _ := url.Parse(myurl)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)

	queryParams := result.Query()
	fmt.Printf("Type of query params are %T\n", queryParams)

	fmt.Println(queryParams)

	for key, val := range queryParams {
		fmt.Printf("%v : %v\n", key, val)
	}

	partsOfUrl := &url.URL{
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hianshu",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}
