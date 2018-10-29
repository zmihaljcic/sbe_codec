//Package codec contains methods for SBE encoding/decoding of messages
package codec

import (
	"bytes"
	"model"
)

//Encoder in an interface that wraps Encode method
//Encode takes a model.Message, encodes and returns encoded bytes
type Encoder interface {
	Encode(msg model.Message) (*bytes.Buffer, error)
}

//Decoder is an interface that wraps Encode method
//Decode decodes input bytes and returns a model.Message
type Decoder interface {
	Decode(b *bytes.Buffer) (model.Message, error)
}

//Codec => encode + decode
type Codec interface {
	Encoder
	Decoder
}

//SBECodec uses (duh) SBE for encoding/decoding
type SBECodec struct {
	M *model.SbeGoMarshaller
}

func (c SBECodec) Encode(msg model.Message) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	err := msg.Encode(c.M, buf, true)
	return buf, err
}

func (c SBECodec) Decode(b *bytes.Buffer) (model.Message, error) {
	msg := model.Message{}
	err := msg.Decode(c.M, b, msg.SbeSchemaVersion(), msg.SbeBlockLength(), true)
	return msg, err
}

//NewSBECodec returns an initialized SBECodec
func NewSBECodec() SBECodec {
	return SBECodec{M: model.NewSbeGoMarshaller()}
}
