package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Create outer slice of length dy
	img := make([][]uint8, dy)

	// Loop for each row
	for y := 0; y < dy; y++ {
		// Create inner slice of length dx
		img[y] = make([]uint8, dx)

		for x := 0; x < dx; x++ {
			// Example pattern: (x + y) / 2
			img[y][x] = uint8((x + y) / 2)
		}
	}

	return img
}

func main() {
	pic.Show(Pic)
}
