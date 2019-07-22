package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func (r13 rot13Reader) Read(b []byte) (int, error) {
    n, err := r13.r.Read(b)

    if err != nil {
        return n, err
    }

    for i, chr := range b {
        switch {
        case 'a' <= chr && (chr <= 'm'):
            b[i] = chr + 13
        case 'n' <= chr && chr <= 'z':
            b[i] = chr - 13
        case 'A' <= chr && chr <= 'M':
            b[i] = chr + 13
        case 'N' <= chr && chr <= 'Z':
            b[i] = chr - 13
        }
    }
    return n, nil
}

func main() {
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

