package controller

import (
	"EchoAPISample/controller"
	"Golang-Echo-MVC-Pattern/model"
	"github.com/labstack/echo"
	"net/http"
)

type ExampleController struct {
	model model.ExampleModel
}

// Interface
type ExampleControllerInterface interface {
	GetPostsController(c echo.Context) error
}

// Get Example Controller
func (ExampleController ExampleController) GetPostsController(c echo.Context) error {
	posts := ExampleController.model.GetPosts()
	res := controller.ResponseGeneric{
		Status:  "Success",
		Message: "Posts Loaded",
		Data:    posts,
	}
	return c.JSON(http.StatusOK, res)
}
