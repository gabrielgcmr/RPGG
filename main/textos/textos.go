package textos

import (
	"aprendendogolang/types"
	"fmt"
)

func AnunciarBatalha(p1, p2 types.Personagem) {
	fmt.Printf("⚔️  Vai começar a batalha entre %s e %s!\n", p1.Nome, p2.Nome)
}
