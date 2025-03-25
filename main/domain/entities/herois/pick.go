package herois

import (
	"main/catalogo/armas"
	"main/catalogo/classes"
	"main/catalogo/raca"
	"main/domain/personagem"
)

var Pick = personagem.NovoPersonagem("Pick", classes.Barbaro, raca.MeioOrc{}, &armas.Arco)
