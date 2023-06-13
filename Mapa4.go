package main

import (
	"fmt"
	"sort"
	"strings"
)

func anagramasPalavras(palavras []string) map[string][]string {
	anagramas := make(map[string][]string)

	for _, palavra := range palavras {
		letras := strings.Split(palavra, "")
		sort.Strings(letras)
		palavraOrdenada := strings.Join(letras, "")
		anagramas[palavraOrdenada] = append(anagramas[palavraOrdenada], palavra)
	}

	return anagramas
}

func main() {
	palavras := []string{"amor", "roma", "carro", "arco", "mar", "ramo"}

	resultado := anagramasPalavras(palavras)
	for chave, grupo := range resultado {
		fmt.Println(chave+":", grupo)
	}
}
