package combate

import (
	"fmt"
	"main/domain/game/dado"
)

func JogadaDeAtaque(atacante Combatente, alvo Combatente) {
	// Rola o d20 para acerto

	if atacante.GetEquipe() == alvo.GetEquipe() {
		fmt.Printf("🛑 %s não pode atacar %s (são aliados!)\n", atacante.GetNome(), alvo.GetNome())
		return
	}

	fmt.Printf("\n⚔️ %s faz uma jogada de ataque em %s!\n", atacante.GetNome(), alvo.GetNome())
	rolagem := dado.Rolar(1, 20)
	// Verifica se acertou
	if rolagem == 20 {
		fmt.Println("💥 CRÍTICO! Acerto automático!")
		dano := atacante.CausarDano() * 2 // Dano dobrado
		fmt.Printf("💢 %s sofre %d de dano!\n", alvo.GetNome(), dano)
		alvo.ReceberDano(dano)
	} else if rolagem == 1 {
		fmt.Println("💩 ERRO CRITICO! Errou feio!")
	} else if rolagem >= alvo.GetCA() {
		fmt.Println("✅ ACERTOU!")
		dano := atacante.CausarDano()
		fmt.Printf("💢 %s sofre %d de dano!\n", alvo.GetNome(), dano)
		alvo.ReceberDano(dano)
	} else {
		fmt.Println("❌ ERROU! O ataque não superou a Classe de Armadura")
	}

}
