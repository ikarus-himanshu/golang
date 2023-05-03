package main

import "fmt"

func main() {
	fmt.Println("Loops in go")

	days := []string{"Sunday", "Monday", "Tuesday", "Wednesday"}
	for d := 0; d < 4; d++ {
		fmt.Println(days[d])
	}

	for i := range days {
		fmt.Println(days[i])
	}

	for index, day := range days {
		fmt.Printf("day %v is %v\n", index, day)
	}

	number := 1
	for number < 10 {

		if number > 5 {
			break
		}

		if number == 2 {
			number++
			continue
		}

		fmt.Printf("Value is %v \n", number)
		number++
	}
}
