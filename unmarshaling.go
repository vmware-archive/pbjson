package pbjson

import (
	"bytes"
	"io"

	"github.com/golang/protobuf/jsonpb"
	"github.com/juju/errors"
)

type decoderOptions struct {
	allowUnknownFields bool
}

type DecoderOption func(*decoderOptions)

func AllowUnknownFields(b bool) DecoderOption {
	return func(o *decoderOptions) {
		o.allowUnknownFields = b
	}
}

type Decoder struct {
	r io.Reader
	jsonpb.Unmarshaler
}

func NewDecoder(r io.Reader, opts ...DecoderOption) *Decoder {
	var o decoderOptions
	for _, opt := range opts {
		opt(&o)
	}
	u := jsonpb.Unmarshaler{AllowUnknownFields: o.allowUnknownFields}
	return &Decoder{r, u}
}

func (d *Decoder) Decode(msg Message) error {
	return errors.Trace(d.Unmarshal(d.r, msg))
}

// UnmarshalString unmarshals a string into a proto message.
func UnmarshalString(str string, msg Message, opts ...DecoderOption) error {
	return errors.Trace(Unmarshal([]byte(str), msg, opts...))
}

// Unmarshal unmarshals a byte array into a proto message.
func Unmarshal(b []byte, msg Message, opts ...DecoderOption) error {
	return errors.Trace(NewDecoder(bytes.NewBuffer(b), opts...).Decode(msg))
}
