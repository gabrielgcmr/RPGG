package combate

import (
	"aprendendogolang/factories"
	"fmt"
)

func EncontrarInimigo(atacante *Participante, ordem []Participante) *Participante {
	for i := range ordem {
		alvo := &ordem[i]

		// Só considera inimigos de tipo diferente e que ainda estão vivos
		if alvo.Tipo != atacante.Tipo && EstaVivo(alvo) {
			return alvo
		}
	}
	return nil // Nenhum inimigo disponível
}

func EstaVivo(p *Participante) bool {
	switch ref := p.Referencia.(type) {
	case *factories.Personagem:
		return ref.PontosDeVida > 0
	case *factories.Monstro:
		return ref.PontosDeVida > 0
	default:
		return false
	}
}

func Ataque(atacante *Participante, alvo *Participante) {
	fmt.Printf("⚔️ %s ataca %s!\n", atacante.Nome, alvo.Nome)
	// Descobre o dano do atacante
	var dano int
	switch a := atacante.Referencia.(type) {
	case *factories.Personagem:
		dano = a.CausarDano()
	case *factories.Monstro:
		dano = a.CausarDano()
	default:
		fmt.Println("❌ Tipo de atacante desconhecido")
		return
	}

	// Aplica o dano no alvo
	switch t := alvo.Referencia.(type) {
	case *factories.Personagem:
		t.PontosDeVida -= dano
		fmt.Printf("💥 %s causou %d de dano em %s (PV restante: %d)\n", atacante.Nome, dano, t.Nome, t.PontosDeVida)
	case *factories.Monstro:
		t.PontosDeVida -= dano
		fmt.Printf("💥 %s causou %d de dano em %s (PV restante: %d)\n", atacante.Nome, dano, t.Nome, t.PontosDeVida)
	default:
		fmt.Println("❌ Tipo de alvo desconhecido")
	}
}
