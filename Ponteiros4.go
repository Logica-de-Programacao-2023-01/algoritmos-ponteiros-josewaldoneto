package main

import (
	"fmt"
	"strconv"
)

// Escreva uma função em Go que receba um ponteiro para uma variável inteira e atualize o valor da variável para a soma dos valores dos seus dois últimos dígitos
// (por exemplo, se o valor inicial da variável for 1234, o novo valor será 3+4=7).

func novoValor(ptr *int) {
	var ultimoN, penultimoN string
	s := strconv.Itoa(*ptr)
	ultimoN += string(s[len(s)-1])
	penultimoN += string(s[len(s)-2])
	uN, _ := strconv.Atoi(ultimoN)
	pN, _ := strconv.Atoi(penultimoN)
	*ptr = uN + pN
}
func main() {
	var inteiro int
	fmt.Println("Digite o valor do inteiro:")
	fmt.Scan(&inteiro)
	novoValor(&inteiro)
	fmt.Println("O novo valor do inteiro é:", inteiro)
}
