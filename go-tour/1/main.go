package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	pow := make([][]uint8, dx)
	for i := range pow {
		pow[i] = make([]uint8, dy)
	}
	return pow
}

func main() {
	pic.Show(Pic)
}
