package domain

import (
	"chat/internal/domain/pubSub"
	"chat/internal/domain/user"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type UseCases struct {
	User userDomain.UserUseCase
	Chat chatDomain.ChatUseCase
}

type Repositories struct {
	User userDomain.UserRepository
	Chat chatDomain.ChatRepository
}

type Config struct {
	Context    context.Context
	Collection *mongo.Collection
}

type DataBase struct {
	ChatCollection *mongo.Collection
	UserCollection *mongo.Collection
	Client         *mongo.Client
}
