package entity

type contextValue string

const (
	JWTCtxKey contextValue = "jwt"
)

type Identity struct {
	ID    int
	Login string
}

func (i Identity) GetID() int {
	return i.ID
}

func (i Identity) GetLogin() string {
	return i.Login
}
