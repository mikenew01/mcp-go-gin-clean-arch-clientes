package entities

import (
	"errors"
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/dtos/input"
	"regexp"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)

type Cliente struct {
	Id    int64  `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func NewCliente() *Cliente {
	return &Cliente{}
}

func FromClienteInput(c *input.ClienteInput) *Cliente {
	cliente := NewCliente()

	if err := cliente.AdicionarNome(c.Nome); err != nil {
		return nil
	}

	if err := cliente.AdicionarEmail(c.Email); err != nil {
		return nil
	}

	return cliente
}

func (c *Cliente) AdicionarEmail(email string) error {

	emailTemporario := c.Email
	c.Email = email

	if err := c.validarEmail(); err != nil {
		c.Email = emailTemporario
		return err
	}

	return nil
}

func (c *Cliente) AdicionarNome(nome string) error {

	nomeTemporario := c.Nome
	c.Nome = nome

	if err := c.validarNome(); err != nil {
		c.Nome = nomeTemporario
		return err
	}

	return nil
}

func (c *Cliente) validarEmail() error {

	if c.Email == "" {
		return errors.New("e-mail está inválido com valor vazio")
	}

	if !emailRegex.MatchString(c.Email) {
		return errors.New("e-mail não está com formato correto")
	}

	return nil

}

func (c *Cliente) validarNome() error {

	if c.Nome == "" {
		return errors.New("nome está inválido com valor vazio")
	}

	return nil

}
