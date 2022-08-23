package chatDomain

import "context"

type Message struct {
	Text string
	ID   uint
}

type Subs struct {
	Name string
	ID   uint
}

type CreateForm struct {
	Context context.Context
	Message Message
}

type ReadForm struct {
	Context context.Context
}

type ChatUseCase interface {
	Sub() string
	Pub() error
}

type ChatRepository interface {
	Create(form CreateForm) error
	Read(form ReadForm) ([]Message, error)
}
