package entities

import (
	"main/catalogo/armas"
	"main/catalogo/classes"
	"main/domain/personagem"
)

var Cactus = personagem.NovoPersonagem(
	"Cactus",
	classes.Barbaro,
	&armas.MachadoSimples,
)
