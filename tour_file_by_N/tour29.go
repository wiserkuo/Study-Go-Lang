package main

import "fmt"

type V struct {
	X, Y int
}

 var (
 	p = V{1, 2}
 	q = &V{1, 2}
 	r = V{X : 3}
 	s = V{}
 )

 func main() {
 	fmt.Println(p, q, r, s)
 }