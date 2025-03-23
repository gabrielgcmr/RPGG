package factories

import "aprendendogolang/types"

func NovoMonstro(
	nome string,
	arma *types.Arma,
	forca, destreza, constituicao, inteligencia, sabedoria, carisma int,
	pontosdevida int,
) types.Monstro {
	atributos := types.Atributos{
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

	return types.Monstro{
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
