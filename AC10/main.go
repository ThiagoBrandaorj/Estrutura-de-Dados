package main

import (
	"fmt"
)

// Função para calcular o MDC de dois números usando o algoritmo de Euclides
func mdc(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	var N, F1, F2 int

	// Lê o número de casos de teste
	fmt.Scanln(&N)

	// Para cada caso de teste
	for i := 0; i < N; i++ {
		// Lê a quantidade de figurinhas de Ricardo e Vicente
		fmt.Scanln(&F1, &F2)

		// Calcula o MDC das quantidades de figurinhas
		maxTroca := mdc(F1, F2)

		// Imprime o tamanho máximo da pilha que pode ser trocada
		fmt.Println(maxTroca)
	}
}
