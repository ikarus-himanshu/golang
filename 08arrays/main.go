package main

import "fmt"

func main() {
	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Banana"

	fmt.Println("My fruit list ", fruitList)
	fmt.Println("My fruit list length ", len(fruitList))

	var vegList = [50]string{"potato", "onion", "mushroom"}

	fmt.Println("My veg list ", vegList)
	fmt.Println("My veg list length ", len(vegList))
}
