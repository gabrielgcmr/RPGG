package combate

import (
	"aprendendogolang/catalogs/dados"
	"aprendendogolang/factories"
	"fmt"
	"sort"
)

func OrdemDeAtaque(personagens []*factories.Personagem, monstro []*factories.Monstro) []Participante {
	fmt.Printf(" Iniciando rodada de combate!\n")
	fmt.Printf(" Definindo iniciativa!\n")
	var participantes []Participante
	//Rolar iniciativas e criar participantes
	for _, p := range personagens {
		fmt.Printf("%s vai rolar 1d20 \n", p.Nome)
		rolagem := dados.RolarDados(1, 20)
		iniciativa := rolagem + p.DES
		fmt.Printf("%s rolou %d + DES %d. Iniciativa = %d \n", p.Nome, rolagem, p.DES, iniciativa)

		participantes = append(participantes, Participante{
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

		participantes = append(participantes, Participante{
			Nome:       m.Nome,
			Iniciativa: iniciativa,
			Tipo:       "Monstro",
			Referencia: m,
		})
	}

	sort.Slice(participantes, func(i, j int) bool {
		return participantes[i].Iniciativa > participantes[j].Iniciativa
	})

	fmt.Println("\n📋 Ordem de Ataque:")
	for i, p := range participantes {
		fmt.Printf("%dº - %s (%s) [Iniciativa: %d]\n", i+1, p.Nome, p.Tipo, p.Iniciativa)
	}
	return participantes
}
