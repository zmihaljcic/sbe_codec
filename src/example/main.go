package main

import (
	"fmt"
	"math/rand"
	"sbeutil/codec"
	"sbeutil/messages"
)

func main() {
	messageFactory := messages.RandomContentMessageFactory{Id: messages.StartingMsgId, Typ: messages.StartingMsgType, Provider: randomBytes}
	message := messageFactory.Create()
	sbeCodec := codec.NewSBECodec()
	buf, err := sbeCodec.Encode(message)

	if err != nil {
		fmt.Println(err)
		return
	}

	messageDecoded, err := sbeCodec.Decode(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Encoded and decoded values are equal? ", (messages.MessagesEqual(message, messageDecoded)))

}

//create 256 random bytes array
func randomBytes() []uint8 {
	bytes := make([]byte, 256)
	rand.Read(bytes)
	return bytes
}
