package main

import "fmt"

func main() {
	defer fmt.Println("Hello")
	fmt.Println("World")
	defer fmt.Println("One")
	fmt.Println("Two")
	defer fmt.Println("Three")
	defer myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
