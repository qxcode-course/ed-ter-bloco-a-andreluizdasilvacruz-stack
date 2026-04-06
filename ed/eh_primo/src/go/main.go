package main

import "fmt"

// x: número que está sendo testado
// div: divisor que está sendo testado
func eh_primo(x int, div int) bool {
	// Casos base para números menores ou iguais a 1 (não são primos)
	if x <= 1 {
		return false
	}
	
	// Ponto de parada de sucesso: se testamos todos os divisores até x e 
	// nenhum deu resto 0, então o número é primo.
	if div == x {
		return true
	}
	
	// Ponto de parada de falha: se encontrar um divisor, não é primo.
	if x%div == 0 {
		return false
	}
	
	// Chamada recursiva testando o próximo divisor
	return eh_primo(x, div+1)
}

func main() {
	var x int
	fmt.Scan(&x)
	if eh_primo(x, 2) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}