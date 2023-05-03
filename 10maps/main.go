package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")
	languages := make(map[string]string)
	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Language map is ", languages)
	fmt.Println("Language short of JS ", languages["JS"])

	delete(languages, "RB")
	fmt.Println("Language map is ", languages)

	//Accessing map using for loop
	for key, value := range languages {
		fmt.Printf("For key %v value is %v\n", key, value)
	}

}
