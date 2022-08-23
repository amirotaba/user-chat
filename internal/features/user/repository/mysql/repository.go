package userRepo

import (
	userDomain "chat/internal/domain/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type mongoRepository struct {
	Collection *mongo.Collection
}

func NewMongoRepository(clc *mongo.Collection) userDomain.UserRepository {
	return &mongoRepository{
		Collection: clc,
	}

}
func (m *mongoRepository) Create(form userDomain.CreateUserForm) error {
	_, err := m.Collection.InsertOne(form.Context, form.User)

	if err != nil {
		return err
	}

	return nil
}

func (m *mongoRepository) Read(form userDomain.ReadUserForm) (userDomain.User, error) {
	cur, err := m.Collection.Find(form.Context, bson.D{})

	if err != nil {
		return userDomain.User{}, err
	}

	defer cur.Close(form.Context)

	var results []userDomain.User

	for cur.Next(form.Context) {

		var result userDomain.User

		if err = cur.Decode(&result); err != nil {
			return userDomain.User{}, err
		}

		results = append(results, result)

	}

	return results[0], nil
}

func (m *mongoRepository) Update(user userDomain.User) error {
	return nil
}

func (m *mongoRepository) Delete(user userDomain.User) error {
	return nil
}
