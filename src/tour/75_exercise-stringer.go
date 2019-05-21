package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
    var s string
    for i, v := range ip {
        if i != 0 {
            s += "." + fmt.Sprintf("%v", v)
        } else {
            s += fmt.Sprintf("%v", v) 
        }
    }
    return s
}

func main() {
    hosts := map[string]IPAddr{
        "loopback": {127, 0, 0, 1},
        "googleDNS": {8, 8, 8, 8},
    }
    for name, ip := range hosts {
        fmt.Printf("%v: %v\n", name, ip)
    }
}
