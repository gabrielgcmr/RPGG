// domain/monstro/methods.go
package monstro

import "main/game/dado"

// --- Getters ---
func (m *Monstro) GetNome() string { return m.Nome }
func (m *Monstro) GetTipo() string { return m.Tipo }
func (m *Monstro) GetCA() int      { return m.CA }
func (m *Monstro) GetHP() int      { return m.HP }
func (m *Monstro) EstaVivo() bool  { return m.HP > 0 }

// --- Modificadores ---
func (m *Monstro) ModFOR() int { return (m.Atributos.FOR - 10) / 2 } // Força
func (m *Monstro) ModDES() int { return (m.Atributos.DES - 10) / 2 } // Destreza
func (m *Monstro) ModCON() int { return (m.Atributos.CON - 10) / 2 } // Constituição
func (m *Monstro) ModINT() int { return (m.Atributos.INT - 10) / 2 } // Inteligência
func (m *Monstro) ModSAB() int { return (m.Atributos.SAB - 10) / 2 } // Sabedoria
func (m *Monstro) ModCAR() int { return (m.Atributos.CAR - 10) / 2 } // Carisma

// --- Combate ---
func (m *Monstro) Iniciativa() int {
	return dado.Rolar(1, 20) + m.ModDES()
}

func (m *Monstro) ReceberDano(dano int) {
	if dano < 0 { // Prevenção contra cura acidental
		dano = 0
	}
	m.HP -= dano
}

func (m *Monstro) Ataque() int {
	if m.Arma != nil {
		return m.Arma.Dano()
	}
	golpe := dado.Rolar(1, 4)
	return golpe + m.ModFOR() // ataque desarmado
}
