package input

type ClienteInput struct {
	Nome  string `json:"nome" validate:"required,min=3,max=100"`
	Email string `json:"email" validate:"required,email, min=3, max=100" `
}
