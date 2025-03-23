package main

import (
	"aprendendogolang/battle"
	"aprendendogolang/entities"
	"aprendendogolang/types"
)

func main() {
	battle.OrdemDeAtaque(
		[]*types.Personagem{&entities.Cactus, &entities.Thor},
		[]*types.Monstro{&entities.AranhaGigante})

}
