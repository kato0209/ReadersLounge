package main

import (
	"backend/controller"
	"backend/controller/openapi"
	"backend/db"
	"backend/repository"
	"backend/router"
	"backend/usecase"
	"fmt"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf(".env is not found")
	}

	e := router.NewRouter()
	_, err = openapi.GetSwagger()
	if err != nil {
		slog.Error("swagger error", "error", err)
	}

	db, err := db.Open()
	if err != nil {
		slog.Error("db connection error", "error", err)
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository)
	server := controller.NewServer(userUsecase)

	openapi.RegisterHandlers(e, server)

	e.Logger.Fatal(e.Start(":8080"))
}
