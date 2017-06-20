package messages

import (
	"github.com/narqo/test-rpc/message"
	"github.com/narqo/test-rpc/runtime"
)

type messageRequestIn struct {
	MessageRequest

	suppliedFields uint32
	emptyFields    uint32
}

var _ runtime.Message = (*messageRequestIn)(nil)

func (msg *messageRequestIn) UnmarshalURL(r *runtime.URLReader) {
	for !r.EOF() {
		switch r.Key() {
		case "id":
			msg.suppliedFields |= 1 << 0
			if r.String() == "" {
				msg.emptyFields |= 1 << 0
				continue
			}
			msg.ID = message.ID(r.Uint64())
		case "offset":
			msg.suppliedFields |= 1 << 1
			if r.String() == "" {
				msg.emptyFields |= 1 << 1
				continue
			}
			msg.Offset = int(r.Int())
		}
	}
}

var _ runtime.URLMarshaler = (*MessageRequest)(nil)

func (req *MessageRequest) MarshalURL(w *runtime.URLWriter) string {
	w.Reset()
	if uint64(req.ID) > 0 {
		w.WriteUint64("id", uint64(req.ID))
	}
	if req.Offset > 0 {
		w.WriteInt("offset", req.Offset)
	}
	return w.String()
}

var _ runtime.URLUnmarshaler = (*MessageRequest)(nil)

func (req *MessageRequest) UnmarshalURL(r *runtime.URLReader) {
	var msg messageRequestIn
	msg.UnmarshalURL(r)
	*req = msg.MessageRequest
}
