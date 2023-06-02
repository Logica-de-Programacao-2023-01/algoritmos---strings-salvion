package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//2. Escreva um programa que receba uma string e remova todos os espaços em branco. Informe ao usuário o resultado.

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("String: ")
	scanner.Scan()
	texto := scanner.Text()
	new_texto := strings.ReplaceAll()
}
