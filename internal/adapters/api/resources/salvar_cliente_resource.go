package resources

import (
	"encoding/json"
	"github.com/maikoncanuto/mcp-service-clientes/internal/adapters/api/handlers"
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/dtos/input"
	"github.com/maikoncanuto/mcp-service-clientes/internal/core/usecases"
	"net/http"
)

type SalvarClienteResource struct {
	usecase usecases.SalvarClienteUseCase
}

func NewSalvarClienteResource(usecase usecases.SalvarClienteUseCase) *SalvarClienteResource {
	return &SalvarClienteResource{
		usecase: usecase,
	}
}

func (resource *SalvarClienteResource) Execute(resp http.ResponseWriter, req *http.Request) {

	var entrada input.ClienteInput

	if err := json.NewDecoder(req.Body).Decode(&entrada); err != nil {
		handlers.NewErrorResponseHandler(err, http.StatusBadRequest).Send(resp)
		return
	}

	defer req.Body.Close()

	output, err := resource.usecase.Execute(req.Context(), &entrada)

	if err != nil {
		handlers.NewErrorResponseHandler(err, http.StatusInternalServerError).Send(resp)
		return
	}

	handlers.NewSuccessResponseHandler(output, http.StatusCreated).Send(resp)
}
