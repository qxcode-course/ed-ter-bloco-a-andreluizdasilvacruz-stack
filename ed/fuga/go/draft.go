package main

import "fmt"

func main() {

    var h, p, f, d int

    fmt.Scan(&h, &p, &f, &d)

    for {

        if f == h {
            fmt.Printf("S\n")
            break
        }

        if f == p {
            fmt.Printf("N\n")
            break
        }

        f = f + (d)

        if f > 15 {
            f = 0
        }

        if f < 0 {
            f = 15
        }
    }
}
