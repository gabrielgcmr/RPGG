// domain/character/methods.go
package personagem

import "main/game/dado"

// --- Getters ---
func (p *Personagem) GetNome() string { return p.Nome }
func (p *Personagem) GetTipo() string { return p.Tipo }
func (p *Personagem) GetRaça() string { return p.Raça }
func (p *Personagem) GetCA() int      { return p.CA }
func (p *Personagem) GetHP() int      { return p.HP }
func (p *Personagem) EstaVivo() bool  { return p.HP > 0 }

// --- Modifiers ---
func (p *Personagem) ModFOR() int { return (p.Atributos.FOR - 10) / 2 }
func (p *Personagem) ModDES() int { return (p.Atributos.DES - 10) / 2 }
func (p *Personagem) ModCON() int { return (p.Atributos.CON - 10) / 2 }
func (p *Personagem) ModINT() int { return (p.Atributos.INT - 10) / 2 }
func (p *Personagem) ModSAB() int { return (p.Atributos.SAB - 10) / 2 }
func (p *Personagem) ModCAR() int { return (p.Atributos.CAR - 10) / 2 }

// --- Combate ---
func (p *Personagem) Iniciativa() int {
	return dado.Rolar(1, 20) + p.ModDES()
}

func (p *Personagem) TomarDano(dano int) {
	if dano < 0 { // Prevenção contra cura acidental
		dano = 0
	}
	p.HP -= dano
}

func (p *Personagem) Atacar() int {
	if p.Arma != nil {
		return p.Arma.Dano()
	}
	soco := dado.Rolar(1, 4)
	return soco + p.ModFOR() // ataque desarmado
}
