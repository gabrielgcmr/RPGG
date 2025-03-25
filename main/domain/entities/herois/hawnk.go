package herois

import (
	"main/catalogo/armas"
	"main/catalogo/classes"
	"main/catalogo/raca"
	"main/domain/personagem"
)

var Hawnk = personagem.NovoPersonagem(
	"Hawnk",
	classes.Bardo,
	raca.Humano{},
	&armas.Adaga,
)
