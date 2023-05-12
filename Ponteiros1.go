package main

import (
	"fmt"
)

// Escreva uma função que receba um ponteiro para um inteiro e um valor inteiro n.
// A função deve atualizar o valor do inteiro para a soma dos n primeiros números naturais.

func atualizar(valor *int, nn int) {
	for i := 0; i <= nn; i++ {
		*valor += i
	}
}
func main() {
	var inteiro, n int
	fmt.Println("Digite o valor de n:")
	fmt.Scan(&n)
	atualizar(&inteiro, n)
	fmt.Println("O novo valor do inteiro é:", inteiro)
}
