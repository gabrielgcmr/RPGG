package main

import (
	"fmt"
	"main/combate"
	"main/entities"
	"main/factories"
)

func main() {
	ordem := combate.Iniciativa(
		[]*factories.Personagem{&entities.Cactus, &entities.Pick},
		[]*factories.Monstro{&entities.AranhaGigante, &entities.Lobo})

	atacante := &ordem[0]
	alvo := combate.EncontrarAlvo(atacante, ordem)
	if alvo != nil {
		combate.Ataque(atacante, alvo)
	} else {
		fmt.Println("⚠️ Nenhum inimigo disponível para atacar.")
	}
}
