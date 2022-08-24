package chat

import (
	"chat/internal/domain/pubSub"
	"chat/internal/middleware/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"net/http"
)

type Handler struct {
	UseCase chatDomain.ChatUseCase
}

func NewHandler(e *echo.Echo, u chatDomain.ChatUseCase) {
	handler := &Handler{
		UseCase: u,
	}

	res := e.Group("user/")
	res.Use(middleware.JWTWithConfig(jwt.Config))
	e.POST("chat/new", handler.NewChat)
	res.POST("chat/", handler.Chat)
	e.GET("chat/receive", handler.Receive)

}

func (m *Handler) NewChat(e echo.Context) error {
	var IDs []primitive.ObjectID
	var ID2 primitive.ObjectID
	if err := e.Bind(&ID2); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	ID := jwt.UserID(e)
	IDs = append(IDs, ID, ID2)

	chat := chatDomain.Chat{
		Subs: IDs,
	}

	form := chatDomain.CreateChat{
		Chat:    chat,
		Context: e.Request().Context(),
	}

	token, err := m.UseCase.NewChat(form)
	if err != nil {
		e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, token)
}

func (m *Handler) Receive(e echo.Context) error {
	chatID := jwt.UserID(e)

	msgs, err := m.UseCase.Receive(chatID)
	if err != nil {
		e.JSON(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, msgs)
}

func (m *Handler) Chat(e echo.Context) error {
	var message chatDomain.Message
	message.UserID = jwt.UserID(e)

	if err := e.Bind(&message); err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	res, err := m.UseCase.Chat(message)
	if err != nil {
		log.Printf("message didn't published %v", err)
	}

	return e.JSON(http.StatusOK, res)
}
