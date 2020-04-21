package estrutura

import "math/rand"

//Jogar - Sortear o Numero
func Jogar(numero int) int {

	numeroSorteado := rand.Intn(50)
	//Contador jogadas
	Contador := 0
	for numeroSorteado != numero {
		numeroSorteado = rand.Intn(50)
		Contador++
	}
	return Contador

}
