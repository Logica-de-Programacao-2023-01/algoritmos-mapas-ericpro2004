package main

import "fmt"

func contarCaracteres(str string) map[rune]int {
	frequencia := make(map[rune]int)
	for _, char := range str {
		frequencia[char]++
	}

	return frequencia
}

func main() {
	str := "Hello, World!"

	resultado := contarCaracteres(str)
	for char, frequencia := range resultado {
		fmt.Printf("%c: %d\n", char, frequencia)
	}
}
