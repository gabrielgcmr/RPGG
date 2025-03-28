package main

import (
	"main/domain/entidades/herois"
	"main/domain/entidades/monstros"
	"main/domain/game/combate"
)

func main() {
	participantes := []combate.Combatente{
		herois.Cactus,
		herois.Bruna,
		monstros.AranhaGigante,
		monstros.Lobo,
	}
	combate.ExecutarCombate(participantes)
}
