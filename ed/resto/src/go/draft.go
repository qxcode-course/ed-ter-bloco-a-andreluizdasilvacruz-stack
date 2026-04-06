package main

import "fmt"

func processa(n int) {
	if n == 0 {
		return
	}
	
	q := n / 2
	r := n % 2
	
	processa(q)
	
	fmt.Printf("%d %d\n", q, r)
}

func main() {
	var n int
	fmt.Scan(&n)
	
	if n == 0 {
		fmt.Println("0 0")
	} else {
		processa(n)
	}
}