package main

import "code.google.com/p/go_tour/pic"

func Pic(dx, dy int) [][]uint8 {
	blue := make([][]uint8, dy)
	for i := range blue {
		blue[i] = make([]uint8, dx)
		for j := range blue[i] {
			blue[i][j] = uint8((i + j) / 2)
		}
	}

}

func main() {
	pic.Show(Pic)
}