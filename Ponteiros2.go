package main

import "fmt"

// Escreva uma função que receba um ponteiro para um inteiro e verifique se esse inteiro é par ou ímpar.
// A função deve atualizar o valor do inteiro para 0 se ele for par ou para 1 se for ímpar.

func verificar(parimpar *int) {
	if *parimpar%2 == 0 {
		*parimpar = 0
	} else {
		*parimpar = 1
	}
}

func main() {
	var inteiro int
	fmt.Println("Digite o valor do inteiro:")
	fmt.Scan(&inteiro)
	verificar(&inteiro)
	fmt.Println("O novo valor do inteiro é:", inteiro)
}
