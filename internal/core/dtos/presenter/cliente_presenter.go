package presenter

import (
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/dtos/output"
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/entities"
)

type ClientePresenter interface {
	Output(cliente *entities.Cliente) (*output.ClienteOutput, error)
}
