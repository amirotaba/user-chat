package user

import (
	userDomain "chat/internal/domain/user"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Handler struct {
	UseCase userDomain.UserUseCase
}

func NewHandler(e *echo.Echo, u userDomain.UserUseCase) {
	handler := &Handler{
		UseCase: u,
	}

	e.POST("signup", handler.SignUp)
	e.POST("signin", handler.SignIn)
}

func (m *Handler) SignIn(e echo.Context) error {
	var loginForm userDomain.LoginForm
	if err := e.Bind(&loginForm); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	u, err := m.UseCase.SignIn(loginForm)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, u)
}

func (m *Handler) SignUp(e echo.Context) error {
	var user userDomain.User
	if err := e.Bind(&user); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	ctx := e.Request().Context()

	form := userDomain.CreateUserForm{
		User:    user,
		Context: ctx,
	}
	u, err := m.UseCase.Create(form)
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusCreated, u)
}
