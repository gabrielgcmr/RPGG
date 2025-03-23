package types

type Monstro struct {
	Nome string

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
