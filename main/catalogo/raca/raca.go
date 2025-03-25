package raca

type Atributos struct {
	FOR int
	DES int
	CON int
	INT int
	SAB int
	CAR int
}

type Raça interface {
	Nome() string
	AplicarBonus(base Atributos) Atributos
}

type Humano struct{}

func (h Humano) Nome() string { return "Humano" }
func (h Humano) AplicarBonus(base Atributos) Atributos {
	base.FOR += 1
	base.DES += 1
	base.CON += 1
	base.INT += 1
	base.SAB += 1
	base.CAR += 1
	return base
}

type Anao struct{}

func (a Anao) Nome() string { return "Anão" }
func (a Anao) AplicarBonus(base Atributos) Atributos {
	base.CON += 2
	return base
}

type Elfo struct{}

func (e Elfo) Nome() string { return "Elfo" }
func (a Elfo) AplicarBonus(base Atributos) Atributos {
	base.DES += 2
	return base
}

type MeioOrc struct{}

func (m MeioOrc) Nome() string { return "Meio-Orc" }
func (m MeioOrc) AplicarBonus(base Atributos) Atributos {
	base.CON += 1
	base.FOR += 2
	return base
}
