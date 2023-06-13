package main

import "fmt"

func calcularSaldo(contas map[string]float64) map[string]float64 {
	total := 0.0
	for _, valor := range contas {
		total += valor
	}

	media := total / float64(len(contas))

	saldo := make(map[string]float64)
	for nome, valor := range contas {
		saldo[nome] = valor - media
	}

	return saldo
}

func main() {
	contas := map[string]float64{
		"João":   100.0,
		"Maria":  50.0,
		"Carlos": 75.0,
		"André":  120.0,
	}

	saldo := calcularSaldo(contas)
	fmt.Println(saldo)
}
