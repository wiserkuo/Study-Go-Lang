package main

import "fmt"

type V struct {
	X int
	Y int
}

func main() {
	v := V{1, 3}
	v.X = 7
	fmt.Println(v.X, v)
}