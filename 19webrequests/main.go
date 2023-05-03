package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://lco.dev"

func main() {
	fmt.Println("Web Requests in golang")
	response, err := http.Get(url)
	checkNilError(err)

	defer response.Body.Close()

	fmt.Printf("Type of response: %T\n", response)

	databytes, err := ioutil.ReadAll(response.Body)

	checkNilError(err)

	content := string(databytes)

	fmt.Println("Content of the request is: ", content)

}
func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}
