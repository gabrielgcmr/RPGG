package types

type Personagem struct {
	Nome        string
	Classe      Classe
	Nivel       int
	Experiencia int

	Atributos Atributos

	FOR int
	DES int
	CON int
	INT int
	SAB int
	CAR int

	ClasseDeArmadura int
	PontosDeVida     int
	Arma             *Arma
}
