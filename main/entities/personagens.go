package entities

import (
	"aprendendogolang/catalogs/arsenal"
	"aprendendogolang/catalogs/classes"
	"aprendendogolang/factories"
)

var (
	Cactus = factories.NovoPersonagem("Cactus", classes.Barbaro, &arsenal.MachadoSimples)
	Pick   = factories.NovoPersonagem("Pick", classes.Barbaro, &arsenal.Arco)
)
