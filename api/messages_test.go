package messages_test

import (
	"reflect"
	"testing"

	"github.com/narqo/test-rpc/api"
	"github.com/narqo/test-rpc/message"
	"github.com/narqo/test-rpc/runtime"
)

func TestMessageRequest_MarshalURL(t *testing.T) {
	testCases := []struct{
		In  messages.MessageRequest
		Out string
	}{
		{
			messages.MessageRequest{},
			"",
		},
		{
			messages.MessageRequest{ID: message.ID(100200), Offset: 2},
			"id=100200&offset=2",
		},
	}

	urlWriter := &runtime.URLWriter{}

	for n, test := range testCases {
		req := test.In.MarshalURL(urlWriter)
		if req != test.Out {
			t.Errorf("case %v,\n got  %+v,\n want %+v", n, req, test.Out)
		}
	}
}

func TestMessageRequest_UnmarshalURL(t *testing.T) {
	testCases := []struct{
		Query  string
		Expect messages.MessageRequest
	}{
		{
			"",
			messages.MessageRequest{},
		},
		{
			"id=100200&offset=2",
			messages.MessageRequest{ID: message.ID(100200), Offset: 2},
		},
	}

	for n, test := range testCases {
		urlReader := runtime.NewURLReader(test.Query)

		var req messages.MessageRequest
		req.UnmarshalURL(urlReader)

		if !reflect.DeepEqual(req, test.Expect) {
			t.Errorf("case %v,\n got  %+v,\n want %+v", n, req, test.Expect)
		}
	}
}
