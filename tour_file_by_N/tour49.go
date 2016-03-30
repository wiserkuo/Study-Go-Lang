package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("mor.")
	case t.Hour() < 17:
		fmt.Println("aft.")
	default:
		fmt.Println("even.")
	}
}