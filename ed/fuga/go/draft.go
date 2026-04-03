package main

import "fmt"

func main() {
    var h, p, f, d int

    fmt.Scan(&h, &p, &f, &d)

    for {
        if f == h {
            fmt.Println("S")
            break
        } 

        if f == p {
            fmt.Println("N")
            break
        }

        f = f + (d)

        if f < 0 {
            f = 15
        }

        if f > 15 {
            f = 0
        }
    }


}
