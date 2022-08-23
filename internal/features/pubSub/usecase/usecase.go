package chatUsecase

import (
	"chat/internal/domain"
	chatDomain "chat/internal/domain/pubSub"
)

type usecase struct {
	ChatRepo chatDomain.ChatRepository
}

func NewUseCase(r domain.Repositories) chatDomain.ChatUseCase {
	return &usecase{
		ChatRepo: r.Chat,
	}
}

func (u usecase) Sub() string {
	//TODO implement me
	panic("implement me")
}

func (u usecase) Pub() error {
	//TODO implement me
	panic("implement me")
}
