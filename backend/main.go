package main

import (
	"backend/controller"
	"backend/controller/openapi"
	"backend/db"
	"backend/repository"
	"backend/router"
	"backend/usecase"
	"backend/validator"

	_ "github.com/lib/pq"
	"golang.org/x/exp/slog"
)

func main() {

	_, err := openapi.GetSwagger()
	if err != nil {
		slog.Error("swagger error", "error", err)
	}

	db, err := db.Open()
	if err != nil {
		slog.Error("db connection error", "error", err)
	}
	defer db.Close()

	userRepository := repository.NewUserRepository(db)
	userValidator := validator.NewUserValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	postRepository := repository.NewPostRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepository)
	server := controller.NewServer(userUsecase, postUsecase)

	e := router.NewRouter(server)

	e.Logger.Fatal(e.Start(":8080"))
}
