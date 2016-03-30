package main

import "fmt"

type V struct {
	Lat, Long float64
}

var m map[string]V

func main() {
	m = make(map[string]V)
	m["TheOne NEO"] = V{23.12, 45.0}
	m["TheOne NEO1"] = V{23.11, 45.08}
	m["TheOne NEO2"] = V{23.11, 45.08}
	m["TheOne NEO3"] = V{23.11, 45.08}
	m["TheOne NEO4"] = V{23.11, 45.08}
	m["TheOne NEO5"] = V{23.11, 45.08}
	m["TheOne NEO6"] = V{23.11, 45.08}
	m["TheOne NEO7"] = V{23.11, 45.08}
	m["TheOne NEO8"] = V{23.11, 45.08}
	m["TheOne NEO9"] = V{23.11, 45.08}
	m["TheOne NEOA"] = V{23.11, 45.08}
	m["TheOne NEOB"] = V{23.11, 45.08}
	fmt.Println(m)
	// fmt.Println(m["TheOne NEO1"])
}