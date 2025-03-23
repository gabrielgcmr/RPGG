package dados

import (
	"fmt"
	"math/rand"
)

func RolarDados(qtd int, lados int) int {
	total := 0
	fmt.Printf("🎲 Rolando %dd%d... ", qtd, lados)

	for i := 0; i < qtd; i++ {
		roll := rand.Intn(lados) + 1
		fmt.Printf(" rolou %d \n", roll)
		total += roll
	}
	return total
}
