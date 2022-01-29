package blog

type Service interface {
	FindContentById(id int) *Blog
}

type Repository interface {
	FindContentById(id int) *Blog
}
