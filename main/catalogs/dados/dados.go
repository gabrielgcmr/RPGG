package dados

import (
	"fmt"
	"math/rand"
)

func RolarDados(qtd int, lados int, bonus int) int {
	total := 0
	fmt.Printf("🎲 Rolando %dd%d", qtd, lados)
	if bonus != 0 {
		fmt.Printf(" + %d", bonus)
	}
	fmt.Print("... ")

	for i := 0; i < qtd; i++ {
		roll := rand.Intn(lados) + 1
		fmt.Printf("[%d] ", roll)
		total += roll
	}

	total += bonus
	fmt.Printf("= %d\n", total)
	return total
}
