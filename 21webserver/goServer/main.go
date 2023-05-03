package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web requests using go")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformPostFormRequest()

}
func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code ", response.StatusCode)
	fmt.Println("Content Length ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	fmt.Println("content ", content)
	// fmt.Println(string(content))

	var responseString strings.Builder
	byteCount, _ := responseString.Write(content)

	fmt.Println("ByteCount is ", byteCount)
	fmt.Println(responseString.String())
}

func PerformPostRequest() {
	const myurl = "http://localhost:8000/post"

	requestData := strings.NewReader(`
	{
		"Name":"Himanshu",
		"Age":"21"
	}
	`)

	response, err := http.Post(myurl, "application/json", requestData)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code ", response.StatusCode)
	fmt.Println("Content Length ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	// fmt.Println("content ", content)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	formdata := url.Values{}
	formdata.Add("firstname", "himanshu")
	formdata.Add("lastname", "kukreja")
	formdata.Add("email", "himanshu@go.dev")

	response, err := http.PostForm(myurl, formdata)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code ", response.StatusCode)
	fmt.Println("Content Length ", response.ContentLength)

	content, _ := ioutil.ReadAll(response.Body)

	// fmt.Println("content ", content)
	fmt.Println(string(content))
}
