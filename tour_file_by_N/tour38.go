package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	blue := make([][]uint8, dy)
	for i := range blue {
		blue[i] = make([]uint8, dx)
		for j := range blue[i] {
			blue[i][j] = uint8((i + j) / 2)
		}
	}
	return blue
}

func main() {
	pic.Show(Pic)
}