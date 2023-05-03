package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{}
	fmt.Printf("Type of fruitList %T\n ", fruitList)

	fruitList = append(fruitList, "Mango", "Banana", "Orange", "Tomato", "Pinrapple")
	fmt.Println("Fruit List ", fruitList)

	fruitList = append(fruitList[1:3])
	fmt.Println("New Fruit List ", fruitList)

	highScores := make([]int, 5)
	highScores[0] = 156
	highScores[1] = 894
	highScores[2] = 326
	highScores[3] = 792
	highScores[4] = 482

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "javacript", "swift", "python", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
