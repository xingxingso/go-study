package main

import (
    "fmt"
)

func countAndSay(n int) string {
	res := "1"

    for n > 1 {
    	length := len(res)
    	tmp := ""
    	for i := 0; i < length; i++ {
    		str := res[i]
    		count := 1
    		for ; i+1 < length && str == res[i+1]; i++ {
    			count++
    		}
    		// fmt.Println(count, fmt.Sprintf("%d", count))
    		tmp += fmt.Sprintf("%d", count)+string(str)
    	}
    	res = tmp
    	n--
    }

    return res
}

func main() {
    res := countAndSay(4)

    fmt.Println(res)
}
