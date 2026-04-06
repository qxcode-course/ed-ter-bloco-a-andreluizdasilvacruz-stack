package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// iguais compara os elementos recursivamente sem usar len() ou laços.
// Utiliza o índice 'i' para andar pelos vetores até encontrar a sentinela.
func iguais(v1, v2 []string, i int) bool {
	// Caso base 1: ambos chegaram na sentinela ao mesmo tempo
	if v1[i] == "END" && v2[i] == "END" {
		return true
	}
	
	// Caso base 2: um chegou na sentinela antes do outro (tamanhos diferentes)
	if v1[i] == "END" || v2[i] == "END" {
		return false
	}
	
	// Caso base 3: os elementos atuais são diferentes
	if v1[i] != v2[i] {
		return false
	}
	
	// Chamada recursiva avançando o índice
	return iguais(v1, v2, i+1)
}

// limpaVetor remove os colchetes, extrai os números e adiciona a sentinela
func limpaVetor(linha string) []string {
	linha = strings.ReplaceAll(linha, "[", " ")
	linha = strings.ReplaceAll(linha, "]", " ")
	
	// Fields já cuida de separar ignorando os espaços em branco extras
	elementos := strings.Fields(linha)
	
	// Adiciona a sentinela no final para sabermos onde o vetor acaba
	elementos = append(elementos, "END")
	
	return elementos
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	scanner.Scan()
	linha1 := scanner.Text()
	
	scanner.Scan()
	linha2 := scanner.Text()
	
	vet1 := limpaVetor(linha1)
	vet2 := limpaVetor(linha2)
	
	if iguais(vet1, vet2, 0) {
		fmt.Println("iguais")
	} else {
		fmt.Println("diferentes")
	}
}