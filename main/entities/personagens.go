package entities

import (
	"aprendendogolang/catalogs/arsenal"
	"aprendendogolang/catalogs/classes"
	"aprendendogolang/factories"
)

var (
	Cactus = factories.NovoPersonagem("Cactus", classes.Barbaro, &arsenal.MachadoSimples)
	Thor   = factories.NovoPersonagem("Thor", classes.Barbaro, &arsenal.Arco)
)
