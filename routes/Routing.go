package routes

import (
	"Golang-Echo-MVC-Pattern/controller"
	"github.com/labstack/echo"
)

type Routing struct {
	example controller.ExampleController
}

func (Routing Routing) GetRoutes() *echo.Echo {
	e := echo.New()
	e.GET("/posts/", Routing.example.GetPostsController)

	return e
}
