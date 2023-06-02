package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//1. Escreva um programa que receba uma string e converta todas as letras minúsculas para maiúsculas.
//Informe ao usuário o resultado.

func main() {
	var texto string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Digite uma string para descapitalizacao: ")
	scanner.Scan()
	texto = scanner.Text()
	minusculo := strings.ToLower(texto)
	fmt.Print(minusculo)
}
