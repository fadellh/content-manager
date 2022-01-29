package blog

type Service interface {
	FindContentById(id int) (*Blog, error)
	FindAllContent() ([]Blog, error)
}

type Repository interface {
	FindContentById(id int) (*Blog, error)
	FindAllContent() ([]Blog, error)
}
