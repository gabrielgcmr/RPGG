// domain/character/type.go
package personagem

import (
	"main/catalogo/armas"
	"main/catalogo/classes"
)

type Personagem struct {
	Nome   string
	Tipo   string // "Herói", "NPC", "Monstro"
	Raça   string
	Classe classes.Classe
	Nivel  int
	XP     int // Experience points

	Atributos struct {
		FOR int
		DES int
		CON int
		INT int
		SAB int
		CAR int
	}

	CA   int // Classe de Armadura
	HP   int // Hit Points
	Arma *armas.Arma
}
