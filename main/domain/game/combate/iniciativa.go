package combate

import (
	"fmt"
	"main/domain/game/dado"

	"sort"
)

func Iniciativa(combatentes []Combatente) []Combatente {
	fmt.Println("\n🎯 Iniciando rodada de combate!")
	fmt.Println("🎲 Rolando iniciativa para os combatentes...")

	type iniciativaCombatente struct {
		combatente Combatente
		valor      int
	}

	var iniciativas []iniciativaCombatente

	// Rolagem para cada combatente
	for _, c := range combatentes {
		fmt.Printf("\n %s vai rolar inciativa... ", c.GetNome())
		rolagem := dado.Rolar(1, 20)
		modificador := c.ModDES()
		total := rolagem + modificador

		fmt.Printf("%s rolou %d (d20) + %d (DES) = %d\n",
			c.GetNome(), rolagem, modificador, total)

		iniciativas = append(iniciativas, iniciativaCombatente{
			combatente: c,
			valor:      total,
		})
	}

	// Ordena do maior para o menor
	sort.Slice(iniciativas, func(i, j int) bool {
		return iniciativas[i].valor > iniciativas[j].valor
	})

	// Cria slice ordenada
	fmt.Println("\n📋 Ordem de Ataque:")
	resultado := make([]Combatente, len(iniciativas))
	for i, item := range iniciativas {
		resultado[i] = item.combatente
		fmt.Printf("%dº: %s (%d)\n", i+1, item.combatente.GetNome(), item.valor)
	}
	return resultado
}
