package routes

import (
	"Golang-Echo-MVC-Pattern/controller"
	"database/sql"
	"github.com/labstack/echo"
)

type Routing struct {
	example controller.ExampleController
}

func (Routing Routing) GetRoutes(db *sql.DB) *echo.Echo {
	e := echo.New()

	e.POST("/posts/", func(c echo.Context) error {
		err := Routing.example.GetPostsController(c, db)
		return err
	})

	return e
}
