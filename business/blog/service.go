package blog

import "content/business"

type service struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &service{
		repository: r,
	}
}

func (s service) FindContentById(id int) (*Blog, error) {

	blog, err := s.repository.FindContentById(id)

	if err != nil {
		return nil, business.ErrDatabase
	}
	if blog.ID == 0 {
		return nil, business.ErrNotFound
	}

	return blog, nil
}
