// combate/interfaces.go
package combate

type Combatente interface {
	GetNome() string
	GetCA() int
	ModDES() int
	CausarDano() int
	ReceberDano(int)
	GetIniciativa() int
	GetEquipe() string
}
