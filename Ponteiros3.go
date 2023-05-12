package main

import (
	"bufio"
	"fmt"
	"os"
)

// Escreva uma função em Go que receba um ponteiro para uma string e atualize o valor da string para seu reverso.

func reverso(str *string) {
	temp := *str
	fraseNova := ""
	for i := len(temp) - 1; i >= 0; i-- {
		fraseNova += string(temp[i])
	}
	*str = fraseNova
}
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite a primeira string.")
	scanner.Scan()
	frase := scanner.Text()
	fmt.Println("A string é: ", frase)
	reverso(&frase)
	fmt.Println("O reverso da string é: ", frase)
}
