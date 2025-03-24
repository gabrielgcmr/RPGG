package combate

type Combatente interface {
	GetNome() string
	GetTipo() string
	GetClasseArmadura() int
	GetPontosDeVida() int
	EstaVivo() bool
	CausarDano() int
	ReceberDano(dano int)
}
