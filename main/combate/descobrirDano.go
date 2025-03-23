package combate

import "aprendendogolang/factories"

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
