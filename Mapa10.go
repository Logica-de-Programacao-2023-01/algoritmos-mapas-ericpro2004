package main

import "fmt"

func contarPares(slice []int) map[[2]int]int {
	frequencia := make(map[[2]int]int)

	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			pair := [2]int{slice[i], slice[j]}
			frequencia[pair]++
		}
	}

	return frequencia
}

func main() {
	slice := []int{1, 2, 3, 4, 5, 1, 2, 3, 4}

	resultado := contarPares(slice)
	for par, frequencia := range resultado {
		fmt.Println(par, ":", frequencia)
	}
}
