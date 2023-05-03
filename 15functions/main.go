package main

import "fmt"

func main() {
	fmt.Println("Functions in go")

	result := adder(3, 5)
	fmt.Println("result of addition", result)

	result, message := proAdder(3, 5, 9, 7, 8, 9, 6)
	fmt.Printf("result of addition is %v and message is %v\n", result, message)
}

func adder(val1 int, val2 int) int {
	return val1 + val2
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi from pro adder"
}
