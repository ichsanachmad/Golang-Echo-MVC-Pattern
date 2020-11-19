package controller

import (
	"Golang-Echo-MVC-Pattern/constant"
	"Golang-Echo-MVC-Pattern/model"
	"Golang-Echo-MVC-Pattern/responsegraph"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"os"
)

type ExampleController struct {
	model model.ExampleModel
}

// generate token example
//func JWTTokenGenerate(data entity.LoginGetData) (entity.Login, error) {
//	currentTime := time.Now()
//	claims := entity.MyCustomClaims{
//		UserID: data.UserID.String,
//		Nama:   data.Name.String,
//		Email:  data.Email.String,
//		RoleID: data.RoleID,
//		StandardClaims: jwt.StandardClaims{
//			ExpiresAt: currentTime.Add(time.Hour * 24).Unix(),
//			IssuedAt:  currentTime.Unix(),
//		},
//	}
//	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
//
//	idToken := entity.Login{}
//	tokenKey := os.Getenv(constant.JWTKey)
//
//	var err error
//	idToken.IDToken, err = token.SignedString([]byte(tokenKey))
//
//	return idToken, err
//}

// Get Example Controller
func (ExampleController ExampleController) GetPostsController(c echo.Context) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	secret := os.Getenv("SECRET_KEY")

	posts := ExampleController.model.GetPosts(secret)
	res := responsegraph.ResponseGenericGet{
		Status:  constant.StatusSuccess,
		Message: constant.MessageSuccess,
		Data:    posts,
	}
	return c.JSON(http.StatusOK, res)
}
