package main

import "fmt"

func main() {
	fmt.Println("Welcome to the class of pointers")

	var ptr *int
	fmt.Println("Value of ptr without initialization is ", ptr)

	myNumber := 23
	ptr = &myNumber

	fmt.Println("Value of just pointer ", ptr)
	fmt.Println("Value at pointer ", *ptr)
	fmt.Println("address of pointer itself ", &ptr)

	*ptr = *ptr + 2
	fmt.Println("New value fo myNumber after pointer operation ", myNumber)

}
