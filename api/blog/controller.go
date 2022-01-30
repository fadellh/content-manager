package blog

import (
	"content/api/blog/request"
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
	b.GET("", handler.FindAllContent)
	b.POST("", handler.InsertContent)
	b.PUT("/:id", handler.UpdateContent)

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

func (ctr *Controller) FindAllContent(c echo.Context) error {
	blogs, err := ctr.service.FindAllContent()

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetAllBlogResponse(blogs)

	return c.JSON(common.NewSuccessResponse(response))

}

func (ctr *Controller) InsertContent(c echo.Context) error {

	insertBlogRequest := new(request.InsertBlogRequest)

	if err := c.Bind(insertBlogRequest); err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	blog, err := ctr.service.InsertContent(request.ToBlogRequest(*insertBlogRequest))

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetBlogResponse(*blog)

	return c.JSON(common.NewSuccessResponse(response))

}

func (ctr *Controller) UpdateContent(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	updateBlogRequest := new(request.UpdateBlogRequest)

	if err := c.Bind(updateBlogRequest); err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	blog, err := ctr.service.UpdateContent(request.ToBlogUpdateRequest(*updateBlogRequest, id))

	if err != nil {
		return c.JSON(common.NewErrorBusinessResponse(err))
	}

	response := response.NewGetBlogResponse(*blog)

	return c.JSON(common.NewSuccessResponse(response))

}
