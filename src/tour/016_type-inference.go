package main

import "fmt"

func main() {
    // v := 42 //int
    // v := 42.0 // float64
    v := 42 + 0.5i // complex128
    fmt.Printf("v is of type %T\n", v)
}
