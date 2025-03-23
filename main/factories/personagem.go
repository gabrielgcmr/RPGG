package factories

import "aprendendogolang/types"

func NovoPersonagem(nome string, classe types.Classe, arma *types.Arma) types.Personagem {
	atributos := types.Atributos{
		Forca:        10,
		Destreza:     10,
		Constituicao: 10,
		Inteligencia: 10,
		Sabedoria:    10,
		Carisma:      10,
	}

	FOR := Modificador(atributos.Forca)
	DES := Modificador(atributos.Destreza)
	CON := Modificador(atributos.Constituicao)
	INT := Modificador(atributos.Inteligencia)
	SAB := Modificador(atributos.Sabedoria)
	CAR := Modificador(atributos.Carisma)

	return types.Personagem{
		Nome:        nome,
		Classe:      classe,
		Nivel:       1,
		Experiencia: 0,

		Atributos: atributos,

		FOR: FOR,
		DES: DES,
		CON: CON,
		INT: INT,
		SAB: SAB,
		CAR: CAR,

		ClasseDeArmadura: atributos.Constituicao + DES,
		PontosDeVida:     classe.VidaBase + CON,
		Arma:             arma,
	}
}
