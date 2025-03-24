package combate

import (
	"fmt"
	"main/catalogs/dados"
	"main/factories"
	"sort"
)

func Iniciativa(personagens []*factories.Personagem, monstro []*factories.Monstro) []Combatente {
	fmt.Printf(" Iniciando rodada de combate!\n")
	fmt.Printf(" Definindo iniciativa!\n")
	var combatente []Combatente
	//Rolar iniciativas e criar participantes
	for _, p := range personagens {
		fmt.Printf("%s vai rolar 1d20 \n", p.Nome)
		rolagem := dados.RolarDados(1, 20)
		iniciativa := rolagem + p.DES
		fmt.Printf("%s rolou %d + DES %d. Iniciativa = %d \n", p.Nome, rolagem, p.DES, iniciativa)

		combatente = append(combatente, Combatente{
			Nome:       p.Nome,
			Iniciativa: iniciativa,
			Tipo:       "Personagem",
			Referencia: p,
		})
	}

	for _, m := range monstro {
		fmt.Printf("%s vai rolar 1d20 \n", m.Nome)
		rolagem := dados.RolarDados(1, 20)
		iniciativa := rolagem + m.DES
		fmt.Printf("%s rolou %d + DES %d. Iniciativa = %d \n", m.Nome, rolagem, m.DES, iniciativa)

		combatente = append(combatente, Combatente{
			Nome:       m.Nome,
			Iniciativa: iniciativa,
			Tipo:       "Monstro",
			Referencia: m,
		})
	}

	sort.Slice(combatente, func(i, j int) bool {
		return combatente[i].Iniciativa > combatente[j].Iniciativa
	})

	fmt.Println("\n📋 Ordem de Ataque:")
	for i, p := range combatente {
		fmt.Printf("%dº - %s (%s) [Iniciativa: %d]\n", i+1, p.Nome, p.Tipo, p.Iniciativa)
	}
	return combatente
}
