package request

import "content/business/blog"

type InsertBlogRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func ToBlogRequest(r InsertBlogRequest) blog.Blog {
	return blog.Blog{
		Title:   r.Title,
		Content: r.Content,
	}
}
