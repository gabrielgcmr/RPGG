package herois

import (
	"main/domain/catalogo/armas"
	"main/domain/catalogo/classes"
	"main/domain/catalogo/raca"
	"main/domain/personagem"
)

var Cactus = personagem.NovoPersonagem(
	"Cactus",
	classes.Barbaro,
	raca.Humano{},
	&armas.MachadoSimples,
)
