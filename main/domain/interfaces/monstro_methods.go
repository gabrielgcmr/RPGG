package interfaces

func (m *Monstro) GetNome() string {
	return m.Nome
}

func (m *Monstro) GetTipo() string {
	return m.Tipo
}

func (m *Monstro) GetClasseArmadura() int {
	return m.ClasseDeArmadura
}

func (m *Monstro) GetPontosDeVida() int {
	return m.PontosDeVida
}

func (m *Monstro) EstaVivo() bool {
	return m.PontosDeVida > 0
}

func (m *Monstro) ReceberDano(dano int) {
	m.PontosDeVida -= dano
}
