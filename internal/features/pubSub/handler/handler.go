package chat

import (
	"chat/internal/domain/pubSub"
	"github.com/labstack/echo/v4"
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

	e.POST("pub", handler.Pub)
	e.GET("sub", handler.Sub)

}

func (m *Handler) Sub(e echo.Context) error {
	//ID := strconv.Itoa(int(jwt.UserID(e)))
	msg := m.UseCase.Sub()
	return e.JSON(http.StatusOK, msg)
}

func (m *Handler) Pub(e echo.Context) error {
	if err := m.UseCase.Pub(); err != nil {
		log.Printf("message didn't published %v", err)
	}
	return e.JSON(http.StatusOK, "message published")
}
