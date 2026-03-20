package main
import "fmt"
func main() {
    var q int
	var d string
	fmt.Scan(&q, &d)

	x := make([]int, q)
	y := make([]int, q)

	for i := 0; i < q; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	// salva posições antigas
	oldX := make([]int, q)
	oldY := make([]int, q)
	copy(oldX, x)
	copy(oldY, y)

	// move a cabeça
	switch d {
	case "L":
		x[0]--
	case "R":
		x[0]++
	case "U":
		y[0]--
	case "D":
		y[0]++
	}

	// move os outros gomos
	for i := 1; i < q; i++ {
		x[i] = oldX[i-1]
		y[i] = oldY[i-1]
	}

	// saída
	for i := 0; i < q; i++ {
		fmt.Printf("%d %d\n", x[i], y[i])
	}

}
