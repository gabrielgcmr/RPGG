package herois

import (
	"main/domain/catalogo/armas"
	"main/domain/catalogo/classes"
	"main/domain/catalogo/raca"
	"main/domain/personagem"
)

var Hawnk = personagem.NovoPersonagem(
	"Hawnk",
	classes.Bardo,
	raca.Humano{},
	&armas.Adaga,
)
