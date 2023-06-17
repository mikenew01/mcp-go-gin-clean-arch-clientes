package output

type ClienteOutput struct {
	Id    int64  `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
