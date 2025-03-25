// domain/personagem/methods.go
package personagem

import (
	"fmt"
	"main/game/dado"
)

// --- Getters ---
func (p *Personagem) GetNome() string   { return p.Nome }
func (p *Personagem) GetEquipe() string { return p.Equipe }
func (p *Personagem) GetRaça() string   { return p.Raça }
func (p *Personagem) GetCA() int        { return p.CA }
func (p *Personagem) GetHP() int        { return p.HP }
func (p *Personagem) EstaVivo() bool    { return p.HP > 0 }

// --- Modifiers ---
func (p *Personagem) ModFOR() int { return (p.Atributos.FOR - 10) / 2 }
func (p *Personagem) ModDES() int { return (p.Atributos.DES - 10) / 2 }
func (p *Personagem) ModCON() int { return (p.Atributos.CON - 10) / 2 }
func (p *Personagem) ModINT() int { return (p.Atributos.INT - 10) / 2 }
func (p *Personagem) ModSAB() int { return (p.Atributos.SAB - 10) / 2 }
func (p *Personagem) ModCAR() int { return (p.Atributos.CAR - 10) / 2 }

// --- Combate ---
func (p *Personagem) GetIniciativa() int {
	return dado.Rolar(1, 20) + p.ModDES()
}
func (p *Personagem) ReceberDano(dano int) {
	if dano < 0 { // Prevenção contra cura acidental
		dano = 0
	}
	p.HP -= dano
	fmt.Printf("💢 %s agora tem %d de HP!\n", p.GetNome(), p.HP)
	if p.HP <= 0 {
		fmt.Println("☠️ MORREU!")
	}
}

func (p *Personagem) CausarDano() int {
	if p.Arma != nil {
		return p.Arma.Dano() + p.ModFOR() // Modificador aplicado
	}
	return dado.Rolar(1, 4) + p.ModFOR() // Ataque desarmado
}
