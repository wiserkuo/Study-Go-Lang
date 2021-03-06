package main

import "fmt"

func fib() func() int {
	a, b := 1, 1
	return func() int {
		a, b = b, a + b
		return b
	}
}

func main() {
	f := fib()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}