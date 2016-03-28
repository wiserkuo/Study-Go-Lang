package main

import "fmt"

type V struct {
	X, Y int
}

func main() {
	var v *V = new(V)
	// v := new(V)
	fmt.Println(v)
	v.X, v.Y = 11, 9
	fmt.Println(v)
}