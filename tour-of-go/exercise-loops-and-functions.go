package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		tmpZ := z - ((z*z - x) / (2 * z))
		if tmpZ == z {
			return z
		}
		z = tmpZ
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
