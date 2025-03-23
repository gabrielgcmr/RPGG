package entities

import (
	"aprendendogolang/catalogs/classes"
	"aprendendogolang/factories"
)

var (
	Cactus = factories.NovoPersonagem("Cactus", classes.Barbaro, &MachadoSimples)
	Thor   = factories.NovoPersonagem("Thor", classes.Barbaro, &Arco)
)
