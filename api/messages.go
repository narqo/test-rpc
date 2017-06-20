package messages

import (
	"context"
	"fmt"

	"github.com/narqo/test-rpc/logger"
	"github.com/narqo/test-rpc/message"
	"github.com/narqo/test-rpc/search"
)

type Messages interface {
	//api:http{"url": "api/v1/messages/message"}
	Message(context.Context, MessageRequest) (MessageResponse, error)

	//api:http{"url": "api/v1/messages/search"}
	Search(context.Context, SearchRequest) (SearchResponse, error)

	//api:http{"url": "api/v1/messages/search/suggest"}
	SearchSuggest(context.Context, SearchSuggestRequest) (SearchSuggestResponse, error)

	//api:http{"url": "api/v1/messages/status"}
	Status(context.Context, StatusRequest) (StatusResponse, error)
}

type MessageRequest struct {
	ID     message.ID `rpc:"required,default:1"`
	Offset int        `rpc:"min:1,max:99"`
}

type MessageResponse struct {
	ID   message.ID
	Text string
}

type SearchRequest struct {
	Query string `rpc:"not_empty"`
}

type SearchResponse struct {
	Results []search.Result
}

type SearchSuggestRequest struct {
	Query string `rpc:"not_empty"`
}

type SearchSuggestResponse struct {
	Results []string
}

type StatusRequest struct {
	ID message.ID
}

type StatusResponse struct {

}


type messagesSvc struct {
	lg logger.Logger
}

func (svc *messagesSvc) Message(ctx context.Context, in MessageRequest) (MessageResponse, error) {
	return MessageResponse{}, fmt.Errorf("not implemented")
}
