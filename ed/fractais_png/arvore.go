package main

import (
	"fmt"
)

// Função recursiva para desenhar a árvore
func desenhaArvore(pen *Pen, tamanho float64) {
	// Ponto de parada: se o galho for muito pequeno, para de desenhar
	if tamanho < 5 {
		return
	}

	// Desenha o galho atual
	pen.Walk(tamanho)

	// --- Galho da Esquerda ---
	pen.Left(30) // vira pra esquerda
	desenhaArvore(pen, tamanho*0.75) // recursão com galho menor
	pen.Right(30) // desfaz a virada para a esquerda

	// --- Galho da Direita ---
	pen.Right(30) // vira pra direita
	desenhaArvore(pen, tamanho*0.75) // recursão com galho menor
	pen.Left(30) // desfaz a virada para a direita

	// Para que a recursão funcione perfeitamente, a caneta tem que voltar 
	// exatamente pro lugar de onde começou a desenhar esse galho
	pen.Up()
	pen.Walk(-tamanho)
	pen.Down()
}

func main() {
	// Cria uma tela 800x800
	pen := NewPen(800, 800) 
	
	// Prepara a caneta no pé da tela, apontando pra cima
	pen.SetPosition(400, 750) 
	pen.SetHeading(90) 
	pen.SetRGB(0, 150, 0) // cor verde
	
	// Começa a desenhar com um tronco de tamanho 120
	desenhaArvore(pen, 120)

	// Salva o resultado
	pen.SavePNG("arvore.png")
	fmt.Println("Árvore gerada em arvore.png")
}