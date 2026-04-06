package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// recursão auxiliar para formatar a string de ida
func tostrRec(vet []int) string {
	if len(vet) == 0 {
		return ""
	}
	if len(vet) == 1 {
		return strconv.Itoa(vet[0])
	}
	return strconv.Itoa(vet[0]) + ", " + tostrRec(vet[1:])
}

func tostr(vet []int) string {
	return "[" + tostrRec(vet) + "]"
}

// recursão auxiliar para formatar a string de volta
func tostrrevRec(vet []int) string {
	if len(vet) == 0 {
		return ""
	}
	if len(vet) == 1 {
		return strconv.Itoa(vet[0])
	}
	// O pulo do gato: processa o resto e anexa o atual no final
	return tostrrevRec(vet[1:]) + ", " + strconv.Itoa(vet[0])
}

func tostrrev(vet []int) string {
	return "[" + tostrrevRec(vet) + "]"
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	// Ponto de parada: se o slice tem 0 ou 1 elemento, já está invertido
	if len(vet) <= 1 {
		return
	}
	// Troca o primeiro com o último
	vet[0], vet[len(vet)-1] = vet[len(vet)-1], vet[0]
	
	// Chama a recursão para o miolo do slice (ignorando o que acabou de trocar)
	reverse(vet[1 : len(vet)-1])
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	return vet[0] + sum(vet[1:])
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0 {
		return 1
	}
	return vet[0] * mult(vet[1:])
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}

	// Função recursiva interna
	var rec func(v []int, start int) (int, int) // retorna (indice_do_menor, valor_do_menor)
	
	rec = func(v []int, start int) (int, int) {
		if start == len(v)-1 {
			return start, v[start]
		}

		// recursão para encontrar o menor no resto do vetor
		indice_resto, valor_resto := rec(v, start+1)

		if v[start] <= valor_resto {
			return start, v[start]
		}
		return indice_resto, valor_resto
	}

	idx_menor, _ := rec(vet, 0)
	return idx_menor
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		
		if len(args) == 0 {
		    continue
		}
		
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			if len(args) > 1 {
			    for _, arg := range args[1:] {
				    if val, err := strconv.Atoi(arg); err == nil {
					    vet = append(vet, val)
				    }
			    }
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}