package combate

import "fmt"

func ExecutarCombate(combatentes []Combatente) {
	ordem := Iniciativa(combatentes)

	fmt.Println("\n⚔️ === COMBATE INICIADO === ⚔️")
	fmt.Println("Ordem de turnos:")
	for i, c := range ordem {
		fmt.Printf("%dº - %s\n", i+1, c.GetNome())
	}

	for rodada := 1; ; rodada++ {
		fmt.Printf("\n=== RODADA %d ===\n", rodada)

		for _, atacante := range ordem {
			if !atacante.EstaVivo() {
				continue
			}

			alvo := EscolherAlvo(atacante, ordem)
			if alvo != nil {
				JogadaDeAtaque(atacante, alvo)
			}

			if CombateAcabou(ordem) {
				fmt.Println("\n⚔️ === COMBATE FINALIZADO === ⚔️")
				return
			}
		}
	}
}

func EscolherAlvo(atacante Combatente, combatentes []Combatente) Combatente {
	// Lógica simplificada: ataca o primeiro inimigo vivo
	for _, c := range combatentes {
		if c.EstaVivo() && c.GetEquipe() != atacante.GetEquipe() {
			return c
		}
	}
	return nil
}

func CombateAcabou(combatentes []Combatente) bool {
	equipesVivas := make(map[string]bool)
	for _, c := range combatentes {
		if c.EstaVivo() {
			equipesVivas[c.GetEquipe()] = true
			if len(equipesVivas) > 1 {
				return false
			}
		}
	}
	return true
}
