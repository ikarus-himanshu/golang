package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 Monday 15:04:05"))

	createdDate := time.Date(2001, time.July, 16, 9, 20, 0, 0, time.Local)
	fmt.Println(createdDate)

	fmt.Println(createdDate.Format("01-02-2006 Monday 15:04:05"))
}
