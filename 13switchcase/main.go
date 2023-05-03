package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Swith case in go")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1")
	case 2:
		fmt.Println("Move 2 spots")
	case 3:
		fmt.Println("Move 3 spots")
		fallthrough
	case 4:
		fmt.Println("Move 4 spots")
	case 5:
		fmt.Println("Move 5 spots")
	case 6:
		fmt.Println("Move 6 spots")
	default:
		fmt.Println("What was that!")
	}

}
