package messages

import (
	"context"

	"github.com/narqo/test-rpc/api/message"
	"github.com/narqo/test-rpc/logger"
)

type Messages interface {
	Message(context.Context, MessageRequest) (MessageResponse, error)
}

type MessageRequest struct {
	ID     message.ID `rpc:"required,default:1"`
	Offset int        `rpc:"min:1,max:99"`
}

type MessageResponse struct {
	ID   message.ID
	Text string
}

type messagesSvc struct {
	lg logger.Logger
}

func (svc *messagesSvc) Message(ctx context.Context, in MessageRequest) (MessageResponse, error) {
	panic("implement me")
}
