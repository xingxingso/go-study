package main

import "fmt"

func main() {
    s := []int{2, 3, 5, 7, 11, 13}
    printSlice(s)

    s = s[:0]
    printSlice(s)

    s = s[:4]
    printSlice(s)

    s = s[2:]
    printSlice(s)

    // s = s[:6]   //panic: runtime error: slice bounds out of range
    // printSlice(s)

    s = s[:4]
    printSlice(s) // len=4 cap=4 [5 7 11 13]

    // s = s[-2:4]
    // printSlice(s) //invalid slice index -2 (index must be non-negative)

    s = s[1:]
    printSlice(s) // len=3 cap=3 [7 11 13]
}

func printSlice(s []int) {
    fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s),s)
}
