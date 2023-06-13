package main

import "fmt"

func somarContagens(contagens []map[string]int) map[string]int {
	soma := make(map[string]int)

	for _, contador := range contagens {
		for palavra, quantidade := range contador {
			soma[palavra] += quantidade
		}
	}

	return soma
}

func main() {
	contagens := []map[string]int{
		{"hello": 2, "world": 1},
		{"hello": 3, "openai": 2},
		{"world": 2, "openai": 1},
	}

	resultado := somarContagens(contagens)
	fmt.Println(resultado)
}
