package herois

import (
	"main/domain/catalogo/armas"
	"main/domain/catalogo/classes"
	"main/domain/catalogo/raca"
	"main/domain/personagem"
)

var Bruna = personagem.NovoPersonagem(
	"Bruna",
	classes.Guerreiro,
	raca.Elfo{},
	&armas.AdagaMagica,
)
