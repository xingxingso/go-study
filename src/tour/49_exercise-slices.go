package main

import "golang.org/x/tour/pic"

import "math"

func Pic(dx, dy int) [][]uint8 {
    x := make([]uint8, dx)
    y := make([][]uint8, dy)

    for i := range y {
        y[i] = x;
        for j := range y[i] {
            y[i][j] = uint8((i+j)/2)
            y[i][j] = uint8(i*j)
            y[i][j] = uint8(math.Pow(float64(i), float64(j)))
            y[i][j] = uint8(float64(i)*math.Log(float64(j))) 
            y[i][j] = uint8(i%(j+1))
            y[i][j] = uint8(float64(i)*math.Sin(float64(j)))
        }
    }
    return y
}

func main() {
    pic.Show(Pic)
}
