package main

import "fmt"

type V struct {
	Lat, Long float64	
}

var m = map[string]V {
	"TheOne" : V{23, 45},
	"air jordan" : V{9, 23},
}

func main() {
	fmt.Println(m);
}