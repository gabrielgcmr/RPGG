package main

import (
	"main/combate"
	"main/domain/entities/herois"
	monstros "main/domain/entities/monsters"
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
