package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
    m := make(map[string]int)
    new_s := strings.Fields(s)
	for _, i := range new_s {
		m[i] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}