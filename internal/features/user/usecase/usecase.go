package userUsecase

import (
	"chat/internal/domain"
	userDomain "chat/internal/domain/user"
	"chat/internal/middleware/jwt"
	"errors"
)

type Usecase struct {
	UserRepo userDomain.UserRepository
}

func NewUseCase(r domain.Repositories) userDomain.UserUseCase {
	return &Usecase{
		UserRepo: r.User,
	}
}

func (a *Usecase) Create(form userDomain.CreateUserForm) (userDomain.UserResponse, error) {
	readForm := userDomain.ReadUserForm{
		UserName: form.User.UserName,
		Context:  form.Context,
	}
	res, _ := a.UserRepo.Read(readForm)
	if res.UserName != "" {
		return userDomain.UserResponse{}, errors.New("this username is taken")
	}

	createForm := userDomain.CreateUserForm{
		User:    form.User,
		Context: form.Context,
	}

	if err := a.UserRepo.Create(createForm); err != nil {
		return userDomain.UserResponse{}, err
	}

	return userDomain.UserResponse{
		UserName: form.User.UserName,
	}, nil
}

func (a *Usecase) SignIn(form userDomain.LoginForm) (userDomain.UserResponse, error) {
	readForm := userDomain.ReadUserForm{
		UserName: form.Username,
		PassWord: form.Password,
	}

	user, err := a.UserRepo.Read(readForm)
	if err != nil {
		return userDomain.UserResponse{}, err
	}

	if form.Password != user.PassWord {
		return userDomain.UserResponse{}, errors.New("wrong password")
	}

	jwtsig, errs := jwt.GenerateTokenUser(user)
	if errs != nil {
		return userDomain.UserResponse{}, errs
	}

	res := userDomain.UserResponse{
		UserName: user.UserName,
		Token:    jwtsig,
	}

	return res, nil
}
