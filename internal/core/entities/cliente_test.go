package entities

import "testing"

func TestNewCliente(t *testing.T) {
	cliente := NewCliente()

	if cliente == nil {
		t.Error("Esperava um novo objeto Cliente, mas recebeu nil")
	}
}

func TestAdicionarEmail(t *testing.T) {
	cliente := NewCliente()
	emailValido := "teste@teste.com"
	emailInvalido := "testetestecom"

	if err := cliente.AdicionarEmail(emailValido); err != nil {
		t.Errorf("Esperava nenhum erro, mas recebeu: %v", err)
	}

	if err := cliente.AdicionarEmail(emailInvalido); err == nil {
		t.Errorf("Esperava um erro, mas não recebeu nenhum")
	}
}

func TestAdicionarEmailVazio(t *testing.T) {
	cliente := NewCliente()
	emailValido := "teste@teste.com"
	emailInvalido := ""

	if err := cliente.AdicionarEmail(emailValido); err != nil {
		t.Errorf("Esperava nenhum erro, mas recebeu: %v", err)
	}

	if err := cliente.AdicionarEmail(emailInvalido); err == nil {
		t.Errorf("Esperava um erro, mas não recebeu nenhum")
	}
}

func TestAdicionarNome(t *testing.T) {
	cliente := NewCliente()
	nomeValido := "Nome Teste"
	nomeInvalido := ""

	if err := cliente.AdicionarNome(nomeValido); err != nil {
		t.Errorf("Esperava nenhum erro, mas recebeu: %v", err)
	}

	if err := cliente.AdicionarNome(nomeInvalido); err == nil {
		t.Errorf("Esperava um erro, mas não recebeu nenhum")
	}
}
