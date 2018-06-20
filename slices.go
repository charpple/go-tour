package main

import "code.google.com/p/go-tour/pic"

func Pic(dx, dy int) [][]uint8 {
    img_func := func(x, y int) uint8 {
        
        return uint8((x+y) / 2)

    }

    img := make([][]uint8, dy)
    for i := range(img) {
        img[i] = make([]uint8, dx)
        for j := range(img[i]) {
            img[i][j] = img_func(i, j)
        }
    }
    return img
}

func main() {
    pic.Show(Pic)
}
