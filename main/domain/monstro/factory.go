// domain/monstro/factory.go
package monstro

import (
	"main/catalogo/armas"
)

// Atributos representa os atributos de um monstro
type Atributos struct {
	FOR int // Força
	DES int // Destreza
	CON int // Constituição
	INT int // Inteligência
	SAB int // Sabedoria
	CAR int // Carisma
}

// NovoPersonagem cria uma nova instância de Personagem com valores padrão
func NovoMonstro(nome string, arma *armas.Arma, atributos Atributos, HP, XP int) *Monstro {

	return &Monstro{
		Nome:      nome,
		Tipo:      "Monstro",
		Atributos: atributos,
		CA:        atributos.CON + ((atributos.DES - 10) / 2),
		HP:        HP,
		XP:        XP,
		Arma:      arma,
	}
}
