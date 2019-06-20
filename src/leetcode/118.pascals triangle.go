package main

import "fmt"

func generate(numRows int) [][]int {
    res := make([][]int, numRows)
    for i:=0; i<numRows; i++ {
        res[i] = make([]int, i+1)
    }
    
    tmp := numRows
    for numRows > 0 {
        res[tmp-numRows][0], res[tmp-numRows][tmp-numRows] = 1, 1
        for i:=1; tmp-numRows-1 > 0 && i<len(res[tmp-numRows-1]); i++ {
            res[tmp-numRows][i] = res[tmp-numRows-1][i-1] + res[tmp-numRows-1][i]
        }
        numRows--
    } 
    
    return res
}

func main() {
    numRows := 5
    res := generate(numRows)

    fmt.Println(res)
}
