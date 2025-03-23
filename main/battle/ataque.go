package battle

import "fmt"

func Atacar(atacante *Participante, alvo *Participante) {
	fmt.Printf("⚔️ %s ataca %s!\n", atacante.Nome, alvo.Nome)
}
