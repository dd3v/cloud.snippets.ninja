package auth

type Service interface{}

type service struct {
}

func NewService() Service {
	return &service{}
}
