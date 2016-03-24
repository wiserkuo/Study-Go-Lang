package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(28, 57))
	sum := func(a, b int) int { return a + b } (3, 4)
	fmt.Println(sum)
}