package main

import (
	"fmt"
	"strings"
)

func contarLetrasPorPalavra(frase string) map[string]map[rune]int {
	resultado := make(map[string]map[rune]int)

	palavras := strings.Fields(frase)

	for _, palavra := range palavras {
		contador := make(map[rune]int)
		for _, letra := range palavra {
			contador[letra]++
		}

		resultado[palavra] = contador
	}

	return resultado
}

func main() {
	frase := "Olá, mundo! Seja bem-vindo."

	resultado := contarLetrasPorPalavra(frase)
	for palavra, contagem := range resultado {
		fmt.Println(palavra+":", contagem)
	}
}
