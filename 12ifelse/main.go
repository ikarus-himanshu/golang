package main

import "fmt"

func main() {
	fmt.Println("if else in golang")
	number := 23
	if number < 10 {
		fmt.Println("Regular user")
	} else if number > 10 {
		fmt.Println("Watch Out")
	} else {
		fmt.Println("Exactly 10")
	}

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num == 3 {
		fmt.Println("Equla to 3")
	} else {
		fmt.Println("Bye bye")
	}
}
