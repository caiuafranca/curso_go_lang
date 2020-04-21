package main

import (
	"CursoUdemy/Iniciante/estrutura"
	"CursoUdemy/Iniciante/matematica"
	"CursoUdemy/Iniciante/welcome"
	"fmt"
)

var (
	endereco string
	//Telefone e uma variavel Publica é preciso comentar
	Telefone   string
	quantidade int
	comprou    bool
	valor      float64
	palavras   rune
)

func main() {

	fmt.Printf("%s\n", welcome.OlaMundo("Caiua França"))
	fmt.Printf("A soma 2 + 2 = %d \n", matematica.Soma(2, 2))

	fmt.Printf("A quantidade de Jogadas foi: %d vezes\n", estrutura.Jogar(30))

}
