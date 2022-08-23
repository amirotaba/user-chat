package utils

import (
	"chat/internal/domain"
	"chat/internal/features/pubSub/handler"
	"chat/internal/features/pubSub/repository"
	"chat/internal/features/pubSub/usecase"
	"chat/internal/features/user/handler/http"
	"chat/internal/features/user/repository/mysql"
	"chat/internal/features/user/usecase"
	"chat/internal/mq/nats"
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func ConnNats() {
	_, err := nats.New()
	if err != nil {
		log.Println("Connecting to message broker failed")
	}
}

func Connection() domain.DataBase {
	dbUser := "root"
	dbPass := "root"
	dbName := "chat_db"
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+dbUser+":"+dbPass+"@localhost:27017/"))
	if err != nil {
		panic(err)
	}
	//defer func() {
	//	if err = client.Disconnect(ctx); err != nil {
	//		panic(err)
	//	}
	//}()
	// Ping the primary
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		panic(err)
	}

	userCollection := client.Database(dbName).Collection("users")
	chatCollection := client.Database(dbName).Collection("chat")

	Db := domain.DataBase{
		UserCollection: userCollection,
		ChatCollection: chatCollection,
		Client:         client,
	}

	fmt.Println("Successfully connected and pinged.")
	return Db
}

//func Migrate(Db *gorm.DB) {
//	_ = Db.AutoMigrate(&pubSubDomain.Message{})
//}

func NewRepository(Db domain.DataBase) domain.Repositories {
	repository := domain.Repositories{
		User: userRepo.NewMongoRepository(Db.UserCollection),
		Chat: chatRepo.NewMysqlRepository(Db.ChatCollection),
	}
	return repository
}

func NewUseCase(repo domain.Repositories) domain.UseCases {
	usecase := domain.UseCases{
		User: userUsecase.NewUseCase(repo),
		Chat: chatUsecase.NewUseCase(repo),
	}
	return usecase
}

func NewHandler(e *echo.Echo, useCase domain.UseCases) {
	chat.NewHandler(e, useCase.Chat)
	user.NewHandler(e, useCase.User)
}
