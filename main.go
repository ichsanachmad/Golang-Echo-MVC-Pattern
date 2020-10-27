package main

import (
	"Golang-Echo-MVC-Pattern/routes"
	"Golang-Echo-MVC-Pattern/settings"
)

// Starting server
func main() {
	echo := routes.Routing.GetRoutes(routes.Routing{}, settings.GetDatabase())
	_ = echo.Start(":1103")
}
