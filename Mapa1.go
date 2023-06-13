package main

import (
	"fmt"
	"strings"
)

func contarPalavras(str string) map[string]int {
	str = strings.ToLower(str)
	str = strings.Map(func(r rune) rune {
		if r >= 'a' && r <= 'z' || r >= '0' && r <= '9' || r == ' ' {
			return r
		}
		return -1
	}, str)
	palavras := strings.Fields(str)
	contagem := make(map[string]int)
	for _, palavra := range palavras {
		contagem[palavra]++
	}

	return contagem
}

func main() {
	str := "Olá, mundo! Olá, Go! Olá, OpenAI!"
	resultado := contarPalavras(str)
	fmt.Println(resultado)
}
