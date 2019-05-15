package main

import (
    "fmt"
    "math"
)

// z -= (z*z - x) / (2*z)
func Sqrt(x float64) float64 {
    z := 1.0;
    for i := 0; i < 10; i++ {
        fmt.Println(z)
        z -= (z*z - x) / (2*z)
    }
    return z
}

func Sqrt2(x float64) float64 {
    z := 1.0;
    tmp := (z*z -x) / (2*z);
    for tmp<(-0.0001) || tmp>0.0001 {
        fmt.Println(z)
        z -= tmp
        tmp = (z*z -x) / (2*z);
    }
    return z
}

func main() {
    fmt.Println(
        Sqrt(2),
        Sqrt2(2),
        math.Sqrt(2),
    )
}
