package combate

type Participante struct {
	Nome       string
	Iniciativa int
	Tipo       string // "Personagem" ou "Monstro"
	Referencia any
}
