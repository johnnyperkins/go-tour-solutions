package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
 x := make([][]uint8, dy)
 z := make([]uint8, dx)

 for i := range x {
 	for ii := range z {
		f := ii
		if f > 128 {
			f = 128 - ii
		} else {
			f = 128 + ii
		}
 		z[ii] = uint8(f)
	 }

	x[i] = z
 }

 return x
}

func main() {
	pic.Show(Pic)
}

