package main

import (
	"CursoUdemy/Iniciante/matematica"
	"CursoUdemy/Iniciante/welcome"
	"testing"
)

func TestOlaMundo(t *testing.T) {

	ola := welcome.Bemvindo
	resposta := "Ola Mundo"

	if resposta != ola {
		t.Error("Aconteceu Algo Errado")
	}

}

func TestOlaMundoPassandoNome(t *testing.T) {

	ola := welcome.OlaMundo("Caiua Franca")
	resposta := "Ola Caiua Franca bem vindo ao sistema!"

	if resposta != ola {
		t.Errorf("Aconteceu Algo Errado, SAIDA: %s |  ESPERADO: %s", ola, resposta)
	}

}

func TestSomarDoisNumeros(t *testing.T) {

	soma := matematica.Soma(2, 2)
	resposta := 4

	if resposta != soma {
		t.Error("Aconteceu Algo Errado")
	}

}

func TestSubtrairDoisNumeros(t *testing.T) {

	sub := matematica.Subtrair(4, 2)
	resultado := 2

	if resultado != sub {
		t.Error("Ocorreu um erro na subtração")
	}

}
