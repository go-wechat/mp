package token

type AccessTokenGetter interface {
	GetToken(code string) (string, error)
	FreshToken() string
}