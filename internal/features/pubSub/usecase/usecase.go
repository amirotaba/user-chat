package chatUsecase

import (
	"chat/internal/domain"
	chatDomain "chat/internal/domain/pubSub"
	userDomain "chat/internal/domain/user"
	"chat/internal/mq/nats"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type usecase struct {
	URepo userDomain.UserRepository
	CRepo chatDomain.ChatRepository
	Mq    nats.Client
}

func NewUseCase(r domain.Repositories) chatDomain.ChatUseCase {
	return &usecase{
		CRepo: r.Chat,
		URepo: r.User,
		Mq:    r.Mq,
	}
}

func (a *usecase) NewChat(form chatDomain.CreateChat) (string, error) {
	if err := a.CRepo.NewChat(form); err != nil {
		return "", err
	}

	//jwtsig, errs := jwt.GenerateTokenChat(form.Chat)
	//if errs != nil {
	//	return "", errs
	//}
	return "jwtsig", nil
}

func (a *usecase) Receive(ID primitive.ObjectID) ([]chatDomain.Message, error) {
	chat := chatDomain.ReadChat{
		ID: ID,
	}
	msgs, err := a.CRepo.ReadChat(chat)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}

func (a *usecase) Chat(message chatDomain.Message) (string, error) {
	if message.Text != "" {
		readForm := userDomain.ReadID{
			ID: message.UserID,
		}

		u, err := a.URepo.ReadID(readForm)
		if err != nil {
			return "", err
		}

		message.Username = u.UserName
		message.Time = time.Now()
		form := chatDomain.CreateMessage{
			Message: message,
		}

		if err := a.CRepo.CreateMessage(form); err != nil {
			return "", err
		}
		a.Mq.Pub(message.Username + message.Text)
		res := a.Mq.Sub()
		return res, nil
	}
	res := a.Mq.Sub()
	return res, nil
}
