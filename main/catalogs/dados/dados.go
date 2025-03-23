package dados

import (
	"math/rand"
)

func D4() int {
	return rolarComPrint(4)
}

func D6() int {
	return rolarComPrint(6)
}

func D8() int {
	return rolarComPrint(8)
}

func D12() int {
	return rolarComPrint(12)
}

func D16() int {
	return rolarComPrint(16)
}

func D20() int {
	return rolarComPrint(20)
}

func Dado(lados int) int {
	return rolarComPrint(lados)
}

// 🔁 Função interna reutilizável
func rolarComPrint(lados int) int {
	result := rand.Intn(lados) + 1
	return result
}
