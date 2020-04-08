package main

import (
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
	resposta := "Ola, Caiua Franca, bem vindo ao sistema"

	if resposta != ola {
		t.Error("Aconteceu Algo Errado")
	}

}
