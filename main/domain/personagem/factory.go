// domain/personagem/factory.go
package personagem

import (
	"main/catalogo/armas"
	"main/catalogo/classes"
	"main/catalogo/raca"
)

const (
	Heroi       = "Herói"
	NPC         = "NPC"
	TipoMonstro = "Monstro"
)

func NovoPersonagem(
	nome string,

	classe classes.Classe,
	raca raca.Raça,
	arma *armas.Arma,
) *Personagem {
	base := struct {
		FOR int
		DES int
		CON int
		INT int
		SAB int
		CAR int
	}{
		FOR: 10,
		DES: 10,
		CON: 10,
		INT: 10,
		SAB: 10,
		CAR: 10,
	}

	atributos := raca.AplicarBonus(base)

	return &Personagem{
		Nome:      nome,
		Equipe:    Heroi,
		Raça:      raca.Nome(),
		Classe:    classe,
		Nivel:     1,
		XP:        0,
		Atributos: atributos,
		CA:        10 + ((atributos.DES - 10) / 2), // Base 10 + DEX mod
		HP:        classe.HPBase + ((atributos.CON - 10) / 2),
		Arma:      arma,
	}
}
