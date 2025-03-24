package dado

import (
	"fmt"
	"math/rand"
)

func Rolar(qtd int, lados int) int {
	total := 0
	fmt.Printf("🎲 Rolando %dd%d... ", qtd, lados)

	for i := 0; i < qtd; i++ {
		result := rand.Intn(lados) + 1
		fmt.Printf(" Resultado: %d \n", result)
		total += result
	}
	return total
}
