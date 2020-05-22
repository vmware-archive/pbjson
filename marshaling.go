package pbjson

import (
	"bytes"
	"io"
	"strings"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
)

type Message proto.Message

type encoderOptions struct {
	indent   string
	defaults bool
}

type EncoderOption func(*encoderOptions)

func WithIndent(indent string) EncoderOption {
	return func(o *encoderOptions) {
		o.indent = indent
	}
}

func WithIndentDepth(indentDepth int) EncoderOption {
	return func(o *encoderOptions) {
		o.indent = strings.Repeat(" ", indentDepth)
	}
}

// WithDefaults ensures all fields are emitted, even when they have the default value (the "zero" value).
func WithDefaults(enable bool) EncoderOption {
	return func(o *encoderOptions) {
		o.defaults = enable
	}
}

type Encoder struct {
	w io.Writer
	jsonpb.Marshaler
}

func NewEncoder(w io.Writer, opts ...EncoderOption) *Encoder {
	var o encoderOptions
	for _, opt := range opts {
		opt(&o)
	}
	m := jsonpb.Marshaler{Indent: o.indent, EmitDefaults: o.defaults}
	return &Encoder{w, m}
}

func (e *Encoder) Encode(msg Message) error {
	return e.Marshal(e.w, msg)
}

func Marshal(msg Message, opts ...EncoderOption) ([]byte, error) {
	var out bytes.Buffer
	if err := NewEncoder(&out, opts...).Encode(msg); err != nil {
		return nil, err
	}
	return out.Bytes(), nil
}

func MarshalToString(msg Message, opts ...EncoderOption) (string, error) {
	var out strings.Builder
	if err := NewEncoder(&out, opts...).Encode(msg); err != nil {
		return "", err
	}
	return out.String(), nil
}
