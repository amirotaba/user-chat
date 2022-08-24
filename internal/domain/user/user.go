package userDomain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	UserName string             `json:"user_name"`
	PassWord string             `json:"pass_word"`
	IsActive bool               `json:"is_active"`
}

type CreateUserForm struct {
	Context context.Context
	User    User
}

type ReadUserForm struct {
	Context  context.Context
	UserName string
	PassWord string
}

type ReadID struct {
	Context context.Context
	ID      primitive.ObjectID
}

type UserForm struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
	IsActive bool   `json:"is_active" gorm:"default:True"`
}

type LoginForm struct {
	Username string `json:"user_name"`
	Password string `json:"pass_word"`
}

type UserResponse struct {
	UserName interface{}
	Token    string
}

type AuthMessage struct {
	Text     string
	UserInfo UserResponse
}

type SignUpMessage struct {
	Text     string
	UserName string
	Email    string
}

type AccountForm struct {
	Uid        uint
	Tp         string
	PageNumber string
}

type UserAccountForm struct {
	Uid      string
	Username string
	ID       string
}

type UserReadForm struct {
	Span   int
	TypeID uint
}

type UserRepository interface {
	Create(form CreateUserForm) error
	Read(form ReadUserForm) (User, error)
	ReadID(form ReadID) (User, error)
	Update(user User) error
	Delete(user User) error
}

type UserUseCase interface {
	Create(form CreateUserForm) (UserResponse, error)
	SignIn(form LoginForm) (UserResponse, error)
}
