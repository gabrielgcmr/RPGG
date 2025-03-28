package herois

import (
	"main/domain/catalogo/armas"
	"main/domain/catalogo/classes"
	"main/domain/catalogo/raca"
	"main/domain/personagem"
)

var Pick = personagem.NovoPersonagem("Pick", classes.Barbaro, raca.MeioOrc{}, &armas.Arco)
