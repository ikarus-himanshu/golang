package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	//there is no heritance in golang; no concept of super ot parent

	himanshu := User{"himanshu kukreja", "himanshu@go.dev", true, 25}
	fmt.Println(himanshu)
	fmt.Printf("himanshu details are: %+v\n", himanshu)
	fmt.Printf("himanshu details are name: %v, email: %v\n", himanshu.Name, himanshu.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
