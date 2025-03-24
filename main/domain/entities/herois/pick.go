package entities

import (
	"main/catalogo/armas"
	"main/catalogo/classes"
	"main/domain/personagem"
)

var Pick = personagem.NovoPersonagem("Pick", classes.Barbaro, &armas.Arco)
