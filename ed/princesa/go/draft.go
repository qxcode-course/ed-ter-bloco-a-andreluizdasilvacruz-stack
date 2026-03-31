package main

import "fmt"

func mostrar_roda(fila []int, pos_espada int) {
    fmt.Print("[ ")
    for i := 0; i < len(fila); i++ {
        fmt.Print(fila[i])
        if i == pos_espada {
            fmt.Print("> ")
        } else {
            fmt.Print(" ")
        }
    }
    fmt.Println("]")
}

func main() {
    var n, e int
    fmt.Scan(&n, &e)

    var fila []int
    for i := 1; i <= n; i++ {
        fila = append(fila, i)
    }

    pos_espada := e - 1

    for len(fila) > 1 {
        mostrar_roda(fila, pos_espada)
        pos_morta := (pos_espada + 1) % len(fila)
        fila = append(fila[:pos_morta], fila[pos_morta+1:]...)
        pos_espada = pos_morta % len(fila)
    }

    mostrar_roda(fila, pos_espada)
}