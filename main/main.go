package main

import (
	"aprendendogolang/battle"
	"aprendendogolang/entities"
	"aprendendogolang/factories"
)

func main() {
	battle.OrdemDeAtaque(
		[]*factories.Personagem{&entities.Cactus, &entities.Thor},
		[]*factories.Monstro{&entities.AranhaGigante})

}
