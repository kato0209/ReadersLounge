package main

import (
	"backend/controllers"
	"backend/controllers/openapi"
	"backend/router"
)

func main() {

	e := router.NewRouter()
	_, err := openapi.GetSwagger()
	if err != nil {
		panic(err)
	}

	server := controllers.NewServer()
	openapi.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":8080"))
}
