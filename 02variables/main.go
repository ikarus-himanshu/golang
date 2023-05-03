package main

import "fmt"

// varibles starting with capital letter are treated as public variable and can be used inside any function

const LoginToken string = "gibbersishvalue"

func main() {

	var username string = "himanshu"
	fmt.Printf("Type of variable: %T \n", username)

	var isLog bool = false
	fmt.Printf("Type of variable: %T \n", isLog)

	var smallval int = 256
	fmt.Printf("Type of variable: %T \n", smallval)

	var floatval float64 = 256.04563615416541653
	fmt.Println(floatval)
	fmt.Printf("Type of variable: %T \n", floatval)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Type of variable: %T \n", anotherVariable)

	var anotherVariable2 string
	fmt.Println(anotherVariable2)
	fmt.Printf("Type of variable: %T \n", anotherVariable2)

	//implicit type
	var anotherVariable3 = "random string"
	fmt.Println(anotherVariable3)
	fmt.Printf("Type of variable: %T \n", anotherVariable3)

	// no var style or called walras syntax
	//This type of declaration can be fine and does'nt require any varible type mentioning just like python does but we can't use (declare) these variables outside the functions

	numberOfUser := 300000
	fmt.Println(numberOfUser)

	// accessing the public variable
	fmt.Println(LoginToken)
}
