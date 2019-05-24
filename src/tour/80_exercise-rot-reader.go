package main

import (
    "io"
    "os"
    "strings"
    // "fmt"
)

type rot13Reader struct {
    r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
    num, err := rot.r.Read(b)

    for i, v := range b {
        if (v >= 'A' && v <= 'M') || (v >= 'a' && v <= 'm') {
            b[i] += 13
        } else if (v > 'M' && v <= 'Z') || (v > 'm' && v <= 'z') {
            b[i] -= 13
        }
    }

    return num, err
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}
