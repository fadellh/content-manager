package blog

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) FindContentById(id int) (*Blog, error) {
	return &Blog{}, nil
}
