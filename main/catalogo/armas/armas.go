package armas

import "main/game/dado"

type Arma struct {
	Nome string
	Dano func() int
}

var (
	Arco = Arma{
		Nome: "Arco Curto",
		Dano: func() int { return dado.Rolar(1, 4) },
	}
	MachadoSimples = Arma{
		Nome: "Machado Simples",
		Dano: func() int { return dado.Rolar(1, 6) },
	}
	Adaga = Arma{
		Nome: "Adaga",
		Dano: func() int { return dado.Rolar(1, 4) },
	}
	Soco = Arma{
		Nome: "Soco",
		Dano: func() int { return dado.Rolar(1, 4) },
	}
	Mordida = Arma{
		Nome: "Mordida",
		Dano: func() int { return dado.Rolar(1, 4) },
	}
)
