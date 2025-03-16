package usecase

type RegisterSupporterOutput struct {
	Id      string `json:"id"`
	Name    string `json:"nome"`
	Email   string `json:"email"`
	Team    string `json:"time"`
	Message string `json:"mensagem"`
}
