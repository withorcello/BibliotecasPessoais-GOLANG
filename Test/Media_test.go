package Test

import "testing"

const msgTest = "Valor esperado %f: \nValor recebido %f:"

func TestMedia(t *testing.T) {
	ValorEsperado := 8.0
	valor := CalculaMedia(8, 8)
	if valor != ValorEsperado {
		t.Errorf(msgTest, ValorEsperado, valor)
	}
}
