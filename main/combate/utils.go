package combate

import "aprendendogolang/factories"

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
