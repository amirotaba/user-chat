package chatRepo

import (
	"chat/internal/domain/pubSub"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	Collection *mongo.Collection
}

func NewMysqlRepository(clc *mongo.Collection) chatDomain.ChatRepository {
	return &mongoRepository{
		Collection: clc,
	}
}

func (m *mongoRepository) NewChat(form chatDomain.CreateChat) error {
	_, err := m.Collection.InsertOne(form.Context, form.Chat)

	if err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository) ReadChat(form chatDomain.ReadChat) ([]chatDomain.Message, error) {
	filter := bson.D{{"chatid", form.ID}}
	cur, err := m.Collection.Find(form.Context, filter)

	if err != nil {
		return nil, err
	}

	defer cur.Close(form.Context)

	var messages []chatDomain.Message

	for cur.Next(form.Context) {

		var message chatDomain.Message

		if err = cur.Decode(&message); err != nil {
			return nil, err
		}

		messages = append(messages, message)

	}

	//res := userDomain.User{}
	//
	//res = userDomain.User{
	//	UserName: products[0].UserName,
	//}

	return messages, nil
}

func (m *mongoRepository) CreateMessage(form chatDomain.CreateMessage) error {
	_, err := m.Collection.InsertOne(form.Context, form.Message)

	if err != nil {
		return err
	}

	return nil
}
