package user

type Service interface {
}

type service struct {
	repository Repository
}

//NewService - ...
func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s *service) GetAll() {

}
