package main

import (
	"fmt"
	"math"
)

// Escreva uma função em Go que receba um ponteiro para uma variável float64 e atualize o valor da variável para a média aritmética
// entre o seu valor atual e o valor da constante Pi.

func media(ptr *float64) {
	pi := math.Pi
	*ptr = (*ptr + pi) / 2
}

func main() {
	var float float64
	fmt.Println("Digite um número:")
	fmt.Scan(&float)
	media(&float)
	fmt.Printf("A média entre o número e pi é, aproximadamente, %.2f\n", float)
}
