package response

import "content/business/blog"

func NewGetAllBlogResponse(b []blog.Blog) []GetBlogResponse {

	var getAllResponse []GetBlogResponse

	for _, val := range b {
		getAllResponse = append(getAllResponse, GetBlogResponse{
			Id:          val.ID,
			Title:       val.Title,
			Content:     val.Content,
			UpdatedAt:   val.UpdatedAt,
			CreatedAt:   val.CreatedAt,
			PublishedAt: val.PublishedAt,
		})
	}

	return getAllResponse
}
