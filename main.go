package main

import "Golang-Echo-MVC-Pattern/routes"

// Starting server
func main() {
	echo := routes.Routing.GetRoutes(routes.Routing{})
	_ = echo.Start(":1337")
}
