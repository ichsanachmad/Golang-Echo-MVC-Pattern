package main

import (
	"Golang-Echo-MVC-Pattern/constant"
	"Golang-Echo-MVC-Pattern/routes"
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

// Starting server
func main() {
	echo := routes.Routing.GetRoutes(routes.Routing{})

	err := godotenv.Load()
	if err != nil {
		println(constant.MessageEnvironment)
	}

	env := os.Getenv(constant.Environment)
	address := os.Getenv(constant.APPHost)
	port := os.Getenv(constant.APPPort)
	host := fmt.Sprint(address, ":", port)

	if env == constant.Local {
		_ = echo.Start(host)
	} else {
		certificate := os.Getenv(constant.CertificateFile)
		key := os.Getenv(constant.KeyFile)
		_ = echo.StartTLS(host, certificate, key)
	}
}
