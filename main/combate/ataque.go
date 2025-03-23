package combate

import "fmt"

func Ataque(atacante *Participante, alvo *Participante) {
	fmt.Printf("⚔️ %s ataca %s!\n", atacante.Nome, alvo.Nome)

	dano := DescobrirDano(atacante)
	if dano == -1 {
		fmt.Println("❌ Falha ao determinar o dano do atacante.")
		return
	}

	AplicarDano(alvo, dano, atacante.Nome)
}
