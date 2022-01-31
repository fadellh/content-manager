package blog

type Service interface {
	FindContentById(id int) (*Blog, error)
	FindAllContent() ([]Blog, error)
	InsertContent(b Blog) (*Blog, error)
	UpdateContent(b Blog) (*Blog, error)
	DeleteContent(id int) (*Blog, error)
}

type Repository interface {
	FindContentById(id int) (*Blog, error)
	FindAllContent() ([]Blog, error)
	InsertContent(b Blog) (*Blog, error)
	UpdateContent(b Blog) error
	DeleteContent(id int) (*Blog, error)
}
