package main

import (
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
	fmt.Printf("A soma 2 + 2 = %d", matematica.Soma(2, 2))

}
