package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("files in golang")
	content := "This is the content of the file that we are going to create \n Gosh it must have lorem ipsum to write the content"
	fileName := "./mytext.txt"

	createFile(content, fileName)
	readFile(fileName)

}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func createFile(content string, filename string) {
	file, err := os.Create(filename)

	checkNilError(err)

	length, err := io.WriteString(file, content)

	checkNilError(err)

	fmt.Println("length is: ", length)

	defer file.Close()

}

func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkNilError(err)
	fmt.Println("Text data inside the file is \n", string(databyte))
}
