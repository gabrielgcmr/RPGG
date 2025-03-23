package main

import (
	"aprendendogolang/combate"
	"aprendendogolang/entities"
	"aprendendogolang/factories"
	"fmt"
)

func main() {
	ordem := combate.OrdemDeAtaque(
		[]*factories.Personagem{&entities.Cactus, &entities.Pick},
		[]*factories.Monstro{&entities.AranhaGigante, &entities.Lobo})

	atacante := &ordem[0]
	alvo := combate.EncontrarInimigo(atacante, ordem)
	if alvo != nil {
		combate.Ataque(atacante, alvo)
	} else {
		fmt.Println("⚠️ Nenhum inimigo disponível para atacar.")
	}
}
