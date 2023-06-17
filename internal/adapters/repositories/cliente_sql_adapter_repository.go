package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/entities"
)

type ClienteSqlAdapterRepository struct {
	db *sql.DB
}

func NewClienteSqlRepository(database *sql.DB) *ClienteSqlAdapterRepository {
	return &ClienteSqlAdapterRepository{db: database}
}

func (repository *ClienteSqlAdapterRepository) Salvar(ctx context.Context, cliente *entities.Cliente) (*entities.Cliente, error) {
	query := `INSERT INTO clientes (nome, email) VALUES ($1, $2) RETURNING id`

	err := repository.db.QueryRowContext(ctx, query, cliente.Nome, cliente.Email).Scan(&cliente.Id)

	if err != nil {
		return nil, fmt.Errorf("error ao inserir cliente: %w", err)
	}

	return cliente, nil
}

func (repository *ClienteSqlAdapterRepository) Atualizar(ctx context.Context, cliente *entities.Cliente) (*entities.Cliente, error) {
	query := `UPDATE clientes SET nome = ?, email = ? WHERE id = ?`

	_, err := repository.db.ExecContext(ctx, query, cliente.Nome, cliente.Email, cliente.Id)

	if err != nil {
		return nil, err
	}

	return cliente, nil
}

func (repository *ClienteSqlAdapterRepository) Deletar(ctx context.Context, id string) error {
	query := `DELETE FROM clientes WHERE id = ?`

	_, err := repository.db.ExecContext(ctx, query, id)

	return err
}

func (repository *ClienteSqlAdapterRepository) ObterPorId(ctx context.Context, id string) (*entities.Cliente, error) {
	query := `SELECT id, nome, email FROM clientes WHERE id = ?`

	row := repository.db.QueryRowContext(ctx, query, id)

	var cliente entities.Cliente

	if err := row.Scan(&cliente.Id, &cliente.Nome, &cliente.Email); err != nil {
		return nil, err
	}

	return &cliente, nil
}

func (repository *ClienteSqlAdapterRepository) ObterTodos(ctx context.Context) ([]*entities.Cliente, error) {
	query := `SELECT id, nome, email FROM clientes`

	rows, err := repository.db.QueryContext(ctx, query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	clientes := make([]*entities.Cliente, 0)

	for rows.Next() {
		var cliente entities.Cliente

		if err := rows.Scan(&cliente.Id, &cliente.Nome, &cliente.Email); err != nil {
			return nil, err
		}

		clientes = append(clientes, &cliente)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return clientes, nil
}
