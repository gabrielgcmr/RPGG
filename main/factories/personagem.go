package factories

import (
	"aprendendogolang/catalogs"
	"aprendendogolang/catalogs/arsenal"
	"aprendendogolang/catalogs/classes"
	"aprendendogolang/catalogs/dados"
	"fmt"
)

type Personagem struct {
	Nome        string
	Classe      classes.Classe
	Nivel       int
	Experiencia int

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

func (p *Personagem) CausarDano() int {
	if p.Arma != nil {
		return p.Arma.Dano()
	}
	soco := dados.RolarDados(1, 4) // ataque desarmado
	fmt.Printf("Personagem %s fez um ataque desarmado com %d de dano.", p.Nome, soco)
	return soco
}

func NovoPersonagem(nome string, classe classes.Classe, arma *arsenal.Arma) Personagem {
	atributos := catalogs.Atributos{
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

	return Personagem{
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
