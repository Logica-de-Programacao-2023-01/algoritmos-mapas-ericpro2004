package main

import "fmt"

func fibonacci(n int) map[int]int {
	sequencia := make(map[int]int)

	a, b := 0, 1
	for i := 0; i <= n; i++ {
		sequencia[i] = a
		a, b = b, a+b
	}

	return sequencia
}

func main() {
	n := 10
	resultado := fibonacci(n)

	for indice, valor := range resultado {
		fmt.Printf("Fibonacci(%d) = %d\n", indice, valor)
	}
}
