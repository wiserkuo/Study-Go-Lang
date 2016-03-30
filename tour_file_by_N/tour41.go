package main

import "fmt"

type V struct {
	Lat, Long float64
}

var m = map[string]V {
	"1st" : {37.1, 50.4},
	"2nd" : {35.4, 32.0},
}

func main() {
	fmt.Println(m)
}