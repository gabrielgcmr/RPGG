package main

import (
	"fmt"
	"main/combate"
	"main/domain/entities/herois"
	monstros "main/domain/entities/monsters"
)

func main() {
	// ordem := combate.Iniciativa([]combate.Combatente{
	// 	herois.Cactus,
	// 	herois.Pick,
	// 	monstros.AranhaGigante,
	// 	monstros.Lobo,
	// })

	fmt.Println(` Porradaria Iniciada! `)

	// Teste 1: Herói ataca monstro (válido)
	fmt.Println("\n=== TESTE 1: Ataque válido ===")
	combate.JogadaDeAtaque(herois.Cactus, monstros.AranhaGigante)

	// Teste 2: Monstro ataca herói (válido)
	fmt.Println("\n=== TESTE 2: Ataque válido ===")
	combate.JogadaDeAtaque(monstros.Ogro, herois.Hawnk)

	// Teste 3: Herói ataca herói (fogo amigo)
	fmt.Println("\n=== TESTE 3: Fogo amigo ===")
	combate.JogadaDeAtaque(herois.Cactus, herois.Pick)

	// Teste 4: Monstro ataca monstro (fogo amigo)
	fmt.Println("\n=== TESTE 4: Fogo amigo ===")
	combate.JogadaDeAtaque(monstros.AranhaGigante, monstros.AranhaGigante)
}
