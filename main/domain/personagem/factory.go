// domain/personagem/factory.go
package personagem

import (
	"main/catalogo/armas"
	"main/catalogo/classes"
)

func NovoPersonagem(
	nome string,
	classe classes.Classe,
	arma *armas.Arma,
) *Personagem {
	atributos := struct {
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

	return &Personagem{
		Nome:      nome,
		Tipo:      "Heroi",
		Raça:      "Humano",
		Classe:    classe,
		Nivel:     1,
		XP:        0,
		Atributos: atributos,
		CA:        10 + ((atributos.DES - 10) / 2), // Base 10 + DEX mod
		HP:        classe.HPBase + ((atributos.CON - 10) / 2),
		Arma:      arma,
	}
}
