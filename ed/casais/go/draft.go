package main

import "fmt"

func main() {
    var n int
    var cont int = 0
    fmt.Scan(&n)

    var vet[50]int

    for i := 0; i < n; i++ {
        fmt.Scan(&vet[i])
    }

    for i := 0; i < n; i++ {
        if vet[i] == 0 {
            continue
        }

        for j := i + 1; j < n; j++ {
            if vet[i] == -vet[j] {
                cont++
                vet[i] = 0
                vet[j] = 0
                break
            }
        }
    }

    fmt.Println(cont)
}
