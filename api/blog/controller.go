package blog

import (
	"content/business/blog"

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
	e.Group("/blog")
}
