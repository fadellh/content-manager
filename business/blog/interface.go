package blog

type Service interface {
	FindContentById(id int) (*Blog, error)
}

type Repository interface {
	FindContentById(id int) (*Blog, error)
}
