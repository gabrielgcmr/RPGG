package battle

import (
	"aprendendogolang/catalogs/dados"
	"aprendendogolang/types"
	"fmt"
	"sort"
)

type Participante struct {
	Nome       string
	Iniciativa int
	Tipo       string // "Personagem" ou "Monstro"
	Referencia any
}

func OrdemDeAtaque(personagens []*types.Personagem, monstro []*types.Monstro) []Participante {
	fmt.Println(" Iniciando rodada de combate!")
	var participantes []Participante
	//Rolar iniciativas e criar participantes
	for _, p := range personagens {
		inic := dados.D20() + p.DES
		fmt.Printf("%s rolou %d \n", p.Nome, inic)

		participantes = append(participantes, Participante{
			Nome:       p.Nome,
			Iniciativa: inic,
			Tipo:       "Personagem",
			Referencia: p,
		})
	}

	for _, m := range monstro {
		inic := dados.D20() + m.DES
		fmt.Printf("%s rolou %d\n", m.Nome, inic)

		participantes = append(participantes, Participante{
			Nome:       m.Nome,
			Iniciativa: inic,
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
