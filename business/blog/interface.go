package blog

type Service interface {
	FindContentById(id int) (*Blog, error)
	FindAllContent() ([]Blog, error)
	InsertContent(b Blog) (*Blog, error)
}

type Repository interface {
	FindContentById(id int) (*Blog, error)
	FindAllContent() ([]Blog, error)
	InsertContent(b Blog) (int, error)
}
