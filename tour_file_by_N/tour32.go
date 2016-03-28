package main

import "fmt"

func main() {
	p := []int {2, 3, 5, 7, 11, 13, 17}
	fmt.Println("pr ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("pr[%d] == %d\n", i, p[i])
	}
}