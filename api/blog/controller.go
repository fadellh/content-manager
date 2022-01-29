package blog

import (
	"content/api/blog/response"
	"content/api/common"
	"content/business/blog"
	"strconv"

	"github.com/labstack/echo"
)

type Controller struct {
	service blog.Service
}

func NewHandler(e *echo.Echo, s blog.Service) {
	handler := &Controller{
		s,
	}
	if handler == nil {
		panic("Controller parameter cannot be nil")
	}
	b := e.Group("/posts")
	b.GET("/:id", handler.FindContentById)
}

func (ctr *Controller) FindContentById(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	blog, err := ctr.service.FindContentById(id)

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetBlogResponse(*blog)

	return c.JSON(common.NewSuccessResponse(response))
}
