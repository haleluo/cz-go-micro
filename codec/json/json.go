// Package json provides a json codec
package json

import (
	"encoding/json"
	"io"

	"c-z.dev/go-micro/codec"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

type Codec struct {
	Conn    io.ReadWriteCloser
	Encoder *json.Encoder
	Decoder *json.Decoder
}

func (c *Codec) ReadHeader(m *codec.Message, t codec.MessageType) error {
	return nil
}

func (c *Codec) ReadBody(b interface{}) error {
	if b == nil {
		return nil
	}
	// Parse the next JSON object from the stream.
	if pb, ok := b.(proto.Message); ok {
		raw := json.RawMessage{}
		if err := c.Decoder.Decode(&raw); err != nil {
			return err
		}
		return protojson.Unmarshal(raw, pb)
	}
	return c.Decoder.Decode(b)
}

func (c *Codec) Write(m *codec.Message, b interface{}) error {
	if b == nil {
		return nil
	}
	return c.Encoder.Encode(b)
}

func (c *Codec) Close() error {
	return c.Conn.Close()
}

func (c *Codec) String() string {
	return "json"
}

func NewCodec(c io.ReadWriteCloser) codec.Codec {
	return &Codec{
		Conn:    c,
		Decoder: json.NewDecoder(c),
		Encoder: json.NewEncoder(c),
	}
}
