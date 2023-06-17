package usecases

import (
	"context"
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/dtos/input"
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/dtos/output"
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/dtos/presenter"
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/entities"
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/repositories"
	"time"
)

type SalvarClienteUseCase interface {
	Execute(ctx context.Context, clienteInput *input.ClienteInput) (*output.ClienteOutput, error)
}

type createClienteUseCase struct {
	repository repositories.ClienteRepository
	presenter  presenter.ClientePresenter
	ctxTimeout time.Duration
}

func NewSalvarClienteUseCase(repository repositories.ClienteRepository, presenter presenter.ClientePresenter, timeout time.Duration) *createClienteUseCase {
	return &createClienteUseCase{
		repository: repository,
		presenter:  presenter,
		ctxTimeout: timeout,
	}
}

func (usecase *createClienteUseCase) Execute(ctx context.Context, clienteInput *input.ClienteInput) (*output.ClienteOutput, error) {

	ctx, cancel := context.WithTimeout(ctx, usecase.ctxTimeout)
	defer cancel()

	cliente := entities.FromClienteInput(clienteInput)

	clienteSalvo, err := usecase.repository.Salvar(ctx, cliente)

	if err != nil {
		return usecase.presenter.Output(&entities.Cliente{})
	}

	return usecase.presenter.Output(clienteSalvo)
}
