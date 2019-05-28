package main

import (
    "fmt"
)

func longestCommonPrefix(strs []string) string {
    s := ""
    for i, v := range strs {
        if i == 0 {
            s = v
            continue
        }
        ilen := len(s)
        jlen := len(v)
        
        for j:=0; j<ilen; j++ {
            if j >= jlen  || s[j] != v[j] {
                s = s[:j]
                break
            }
        }   
    }    
    return s
}

func main() {
    strs := []string{"flower","flow","flight"}
    // strs := []string{"dog","racecar","car"}
    // strs := []string{"aa","a"}
    
    res := longestCommonPrefix(strs)

    fmt.Println(res)
}
