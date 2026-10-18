package auth

type Service struct {
	repository Repository
}

type Repository interface {

}

func New (repository Repository) Service {
	return Service{
		repository: repository,
	}
}