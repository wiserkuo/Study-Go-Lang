package main

import "fmt"

func main() {
	sum := 1
	a := 1
	for a > 0 && sum != 0 {
		sum += sum
		fmt.Println(sum)
	}
}