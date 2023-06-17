package repositories

import (
	"context"
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/entities"
)

type ClienteRepository interface {
	Salvar(ctx context.Context, cliente *entities.Cliente) (*entities.Cliente, error)
	Atualizar(ctx context.Context, cliente *entities.Cliente) (*entities.Cliente, error)
	Deletar(ctx context.Context, id string) error
	ObterPorId(ctx context.Context, id string) (*entities.Cliente, error)
	ObterTodos(ctx context.Context) ([]*entities.Cliente, error)
}
