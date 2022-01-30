package request

import "content/business/blog"

type UpdateBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func ToBlogUpdateRequest(r UpdateBlogRequest, id int) blog.Blog {
	return blog.Blog{
		ID:      id,
		Title:   r.Title,
		Content: r.Content,
	}
}
