package factories

import (
	"aprendendogolang/catalogs"
	"aprendendogolang/catalogs/arsenal"
)

type Monstro struct {
	Nome string

	Atributos catalogs.Atributos

	FOR int
	DES int
	CON int
	INT int
	SAB int
	CAR int

	ClasseDeArmadura int
	PontosDeVida     int
	Arma             *arsenal.Arma
}

func NovoMonstro(
	nome string,
	arma *arsenal.Arma,
	forca, destreza, constituicao, inteligencia, sabedoria, carisma int,
	pontosdevida int,
) Monstro {
	atributos := catalogs.Atributos{
		Forca:        forca,
		Destreza:     destreza,
		Constituicao: constituicao,
		Inteligencia: inteligencia,
		Sabedoria:    sabedoria,
		Carisma:      carisma,
	}
	FOR := Modificador(forca)
	DES := Modificador(destreza)
	CON := Modificador(constituicao)
	INT := Modificador(inteligencia)
	SAB := Modificador(sabedoria)
	CAR := Modificador(carisma)

	return Monstro{
		Nome:      nome,
		Atributos: atributos,

		FOR: FOR,
		DES: DES,
		CON: CON,
		INT: INT,
		SAB: SAB,
		CAR: CAR,

		ClasseDeArmadura: constituicao + DES,
		PontosDeVida:     pontosdevida,
		Arma:             arma,
	}
}
