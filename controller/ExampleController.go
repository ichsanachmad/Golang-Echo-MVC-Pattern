package controller

import (
	"Golang-Echo-MVC-Pattern/constant"
	"Golang-Echo-MVC-Pattern/model"
	"Golang-Echo-MVC-Pattern/responsegraph"
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
)

type ExampleController struct {
	model model.ExampleModel
}

// Get Example Controller
func (ExampleController ExampleController) GetPostsController(c echo.Context, db *sql.DB) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := os.Getenv("SECRET_KEY")

	posts := ExampleController.model.GetPosts(secret, db)
	res := responsegraph.ResponseGenericGet{
		Status:  constant.StatusSuccess,
		Message: constant.MessageSuccess,
		Data:    posts,
	}
	return c.JSON(http.StatusOK, res)
}
