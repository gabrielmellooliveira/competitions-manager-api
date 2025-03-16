package interfaces

type Authenticator interface {
	GenerateToken(value string) (string, error)
	ValidateToken(token string) (string, error)
}
