//Package messages provides basic utilities for creating model.Message instances
package messages

import (
	"bytes"
	"model"
)

//Some message id/type constants
const (
	StartingMsgId   = uint64(1)
	StartingMsgType = uint8(1)
)

//RandBytesProvider generates an array of random
type RandBytesProvider func() []uint8

//RandomContentMessageFactory generates messages with contstant id and type, but with random bytes
//for content. Random bytes are generated using given provider
type RandomContentMessageFactory struct {
	Provider RandBytesProvider
	Id       uint64
	Typ      uint8
}

//MessageFactory provides handy methods for creating model.Message
type MessageFactory interface {
	Create() model.Message
}

func (factory *RandomContentMessageFactory) Create() model.Message {
	return model.Message{MsgId: factory.Id, MsgType: factory.Typ, Content: factory.Provider()}
}

//MessagesEqual checks if 2 messages are equal
func MessagesEqual(msg1, msg2 model.Message) bool {
	return msg1.MsgId == msg2.MsgId && msg1.MsgType == msg2.MsgType && bytes.Equal(msg1.Content, msg2.Content)
}
