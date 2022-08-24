package chatDomain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Chat struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Subs []primitive.ObjectID
	Time time.Time
}

type Message struct {
	UserID   primitive.ObjectID
	Username string    `json:"username"`
	Text     string    `json:"text"`
	Time     time.Time `json:"time"`
}

type Subs struct {
	Name string
	ID   uint
}

type CreateChat struct {
	Chat    Chat
	Context context.Context
}

type ReadChat struct {
	ID      primitive.ObjectID
	Context context.Context
}

type CreateMessage struct {
	Context context.Context
	Message Message
}

type UserNewChat struct {
	Subs []primitive.ObjectID
}

type ReadChatForm struct {
	ID        primitive.ObjectID
	ChatToken primitive.ObjectID
}

type ChatUseCase interface {
	NewChat(form CreateChat) (string, error)
	Receive(ID primitive.ObjectID) ([]Message, error)
	Chat(Message) (string, error)
}

type ChatRepository interface {
	NewChat(form CreateChat) error
	CreateMessage(form CreateMessage) error
	ReadChat(form ReadChat) ([]Message, error)
}
