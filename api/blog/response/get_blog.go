package response

import "content/business/blog"

type GetBlogResponse struct {
	Id          int    `json:"id"`
	Title       string `json:"title"`
	Content     string `json:"content"`
	UpdatedAt   string `json:"updated_at"`
	CreatedAt   string `json:"created_at"`
	PublishedAt string `json:"published_at"`
}

func NewGetBlogResponse(b blog.Blog) *GetBlogResponse {

	blogResponse := GetBlogResponse{
		Id:          b.ID,
		Title:       b.Title,
		Content:     b.Content,
		PublishedAt: b.PublishedAt,
		UpdatedAt:   b.UpdatedAt,
		CreatedAt:   b.CreatedAt,
	}

	return &blogResponse
}
