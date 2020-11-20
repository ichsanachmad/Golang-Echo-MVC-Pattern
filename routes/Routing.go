package routes

import (
	"Golang-Echo-MVC-Pattern/controller"
	"Golang-Echo-MVC-Pattern/utils"
	"github.com/labstack/echo"
)

type Routing struct {
	example controller.ExampleController
}

func (Routing Routing) GetRoutes() *echo.Echo {
	e := echo.New()

	e.GET("/posts/", Routing.example.GetPostsController)
	e.POST("/posts/", Routing.example.GetPostsController, utils.Middleware)

	return e
}
