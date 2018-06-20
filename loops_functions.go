package main

import (
	"fmt"
	"math"
)

const error = 1e-8

func Sqrt(x float64) float64 {

    z := x
	var y=0
    for ; y<10; y++ {
        nz := z - ((z*z - x) / (2*z))
        if math.Abs(nz - z) < error {
            return nz
        }
        z = nz
    }
	return z
}

func main() {

	fmt.Println("Formula result:", Sqrt(4))

	fmt.Println("Square root function result: ",math.Sqrt(4))
}
