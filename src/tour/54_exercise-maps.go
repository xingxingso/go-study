package main

import (
    "golang.org/x/tour/wc"
    "strings"
    // "fmt"
)

func WordCount(s string) map[string]int {
    stringArr := strings.Fields(s)

    strCount := make(map[string]int)

    for _, v := range stringArr {
        if _, ok := strCount[v]; ok {
            strCount[v]++
        } else {
            strCount[v] = 1
        }
    }
    // fmt.Println(stringArr, strCount)

    // return map[string]int{"x": 1}
    return strCount
}

func main() {
    // WordCount("This is a good test . is it? test test")
    wc.Test(WordCount)
}
