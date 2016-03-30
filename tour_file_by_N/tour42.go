package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["A"] = 46
	fmt.Println(m["A"])

	m["A"] = 99
	fmt.Println(m["A"])

	delete(m, "A")
	fmt.Println(m["A"])

	v, ok := m["A"]
	fmt.Println(v, ok)
}