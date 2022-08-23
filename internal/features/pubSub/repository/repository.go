package chatRepo

import (
	"chat/internal/domain/pubSub"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mysqlUserRepository struct {
	Clc *mongo.Collection
}

func NewMysqlRepository(cnf *mongo.Collection) chatDomain.ChatRepository {
	return &mysqlUserRepository{
		Clc: cnf,
	}
}

func (m *mysqlUserRepository) Create(form chatDomain.CreateForm) error {
	_, err := m.Clc.InsertOne(form.Context, form.Message)
	if err != nil {
		return err
	}
	return nil
}

func (m *mysqlUserRepository) Read(form chatDomain.ReadForm) ([]chatDomain.Message, error) {
	cur, err := m.Clc.Find(form.Context, bson.D{})

	if err != nil {
		return nil, err
	}

	defer cur.Close(form.Context)

	var products []chatDomain.Message

	for cur.Next(form.Context) {

		var product chatDomain.Message

		if err = cur.Decode(&product); err != nil {
			return nil, err
		}

		products = append(products, product)

	}

	//res := userDomain.User{}
	//
	//res = userDomain.User{
	//	UserName: products[0].UserName,
	//}

	return products, nil
}
