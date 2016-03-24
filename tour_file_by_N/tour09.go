package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	// a, b := swap("1st string", "2nd string")
	var a = "1st string"
	var b = "2nd string"
	c, d := swap(a, b)
	fmt.Println(c, d)
}