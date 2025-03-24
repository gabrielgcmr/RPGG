package combate

import (
	"aprendendogolang/factories"
	"fmt"
)

func Ataque(atacante *Participante, alvo *Participante) {
	fmt.Printf("⚔️ %s ataca %s!\n", atacante.Nome, alvo.Nome)

	dano := DescobrirDano(atacante)
	if dano == -1 {
		fmt.Println("❌ Falha ao determinar o dano do atacante.")
		return
	}

	AplicarDano(alvo, dano, atacante.Nome)
}

func DescobrirDano(atacante *Participante) int {
	switch a := atacante.Referencia.(type) {
	case *factories.Personagem:
		return a.CausarDano()
	case *factories.Monstro:
		return a.CausarDano()
	default:
		return -1
	}
}

func AplicarDano(alvo *Participante, dano int, nomeDoAtacante string) {
	switch t := alvo.Referencia.(type) {
	case *factories.Personagem:
		t.PontosDeVida -= dano
		fmt.Printf("💥 %s causou %d de dano em %s (PV restante: %d)\n", nomeDoAtacante, dano, t.Nome, t.PontosDeVida)
	case *factories.Monstro:
		t.PontosDeVida -= dano
		fmt.Printf("💥 %s causou %d de dano em %s (PV restante: %d)\n", nomeDoAtacante, dano, t.Nome, t.PontosDeVida)
	default:
		fmt.Println("❌ Tipo de alvo desconhecido")
	}
}
