package main

import (
	"backend/controller"
	"backend/controller/openapi"
	"backend/db"
	"backend/initialize"
	"backend/models/chat"
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

	initialize.Init(db)

	userRepository := repository.NewUserRepository(db)
	userValidator := validator.NewUserValidator()
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)

	postRepository := repository.NewPostRepository(db)
	postUsecase := usecase.NewPostUsecase(postRepository)

	bookRepository := repository.NewBookRepository(db)
	bookUsecase := usecase.NewBookUsecase(bookRepository)

	chatRepository := repository.NewChatRepository(db)
	chatUsecase := usecase.NewChatUsecase(chatRepository)

	connectionRepository := repository.NewConnectionRepository(db)
	connectionUsecase := usecase.NewConnectionUsecase(connectionRepository)

	hub := chat.NewHub()
	server := controller.NewServer(userUsecase, postUsecase, bookUsecase, chatUsecase, connectionUsecase, *hub)

	go server.RunLoop(hub)

	e := router.NewRouter(server)

	e.Logger.Fatal(e.Start(":8080"))
}
