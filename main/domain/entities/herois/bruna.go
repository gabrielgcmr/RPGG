package herois

import (
	"main/catalogo/armas"
	"main/catalogo/classes"
	"main/catalogo/raca"
	"main/domain/personagem"
)

var Bruna = personagem.NovoPersonagem(
	"Bruna",
	classes.Guerreiro,
	raca.Elfo{},
	&armas.AdagaMagica,
)
