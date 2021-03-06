package blog

import (
	"content/business"
)

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

func (s service) FindAllContent() ([]Blog, error) {

	var blogs []Blog

	blogs, err := s.repository.FindAllContent()

	if err != nil {
		return nil, err
	}

	if len(blogs) == 0 {
		return nil, business.ErrNotFound
	}

	return blogs, nil
}

func (s service) InsertContent(b Blog) (*Blog, error) {

	blog, err := s.repository.InsertContent(b)

	if err != nil {
		return nil, business.ErrDatabase
	}

	return blog, nil
}

func (s service) UpdateContent(b Blog) (*Blog, error) {

	err := s.repository.UpdateContent(b)

	if err != nil {
		return nil, business.ErrDatabase
	}

	blog, _ := s.FindContentById(b.ID)

	return blog, nil
}

func (s service) DeleteContent(id int) (*Blog, error) {

	blog, err := s.repository.DeleteContent(id)

	if err != nil {
		return nil, business.ErrDatabase
	}

	if blog.ID == 0 {
		return nil, business.ErrNotFound
	}

	return blog, nil
}
