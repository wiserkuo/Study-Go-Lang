package main

import "fmt"

type V struct {
	X int
	Y int
}

func main() {
	p := V{1, 2}
	q := &p
	q.X = 1e9
	fmt.Println(p, q)
}