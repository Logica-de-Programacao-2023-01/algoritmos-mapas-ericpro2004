package main

import "fmt"

func somarValores(mapa map[string]int) int {
	soma := 0

	for _, valor := range mapa {
		soma += valor
	}

	return soma
}

func main() {
	mapa := map[string]int{
		"chave1": 10,
		"chave2": 20,
		"chave3": 30,
	}

	resultado := somarValores(mapa)
	fmt.Println("Soma:", resultado)
}
