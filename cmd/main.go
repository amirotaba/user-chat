package main

import (
	"chat/internal/utils"
	"context"
	"github.com/labstack/echo/v4"
	"log"
)

func main() {
	//connect to nats
	utils.ConnNats()

	//connect to database
	Db := utils.Connection()
	//
	//migrate tables
	//utils.Migrate(Db)
	//
	//start echo
	e := echo.New()
	//
	////get repositories
	repos := utils.NewRepository(Db)
	if err := Db.Client.Connect(context.Background()); err != nil {
		log.Println(err)
	}
	//
	////get UseCases
	useCase := utils.NewUseCase(repos)
	//
	//register features
	utils.NewHandler(e, useCase)
	//
	////route
	e.Logger.Fatal(e.Start(":4000"))
}
