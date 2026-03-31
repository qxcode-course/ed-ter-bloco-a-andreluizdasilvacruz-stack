
package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	fila := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	var m int
	fmt.Scan(&m)

	sairam := make(map[int]bool)
	for i := 0; i < m; i++ {
		var id int
		fmt.Scan(&id)
		sairam[id] = true
	}

	for i := 0; i < n; i++ {
		if !sairam[fila[i]] {
			fmt.Print(fila[i], " ")
		}
	}
	fmt.Println()
}