package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Sat?")
	today := time.Now().Weekday()

	fmt.Println(time.Now())
	fmt.Println(today)
	fmt.Println(time.Saturday)
	
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("See ur calendar.")
	}
}