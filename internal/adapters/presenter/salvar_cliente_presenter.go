package presenter

import (
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/dtos/output"
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/dtos/presenter"
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/entities"
)

type salvarClientePresenterInterno struct{}

var _ presenter.ClientePresenter = (*salvarClientePresenterInterno)(nil)

func NewSalvarClientePresenter() *salvarClientePresenterInterno {
	return &salvarClientePresenterInterno{}
}

func (presenter *salvarClientePresenterInterno) Output(cliente *entities.Cliente) (*output.ClienteOutput, error) {
	return &output.ClienteOutput{
		Id:    cliente.Id,
		Nome:  cliente.Nome,
		Email: cliente.Email,
	}, nil
}
