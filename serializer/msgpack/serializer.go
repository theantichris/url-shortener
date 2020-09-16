package msgpack

import (
	"github.com/pkg/errors"
	"github.com/theantichris/url-shortener/shortener"
	"github.com/vmihailenco/msgpack"
)

// Redirect holds a JSON implementation of shortener.Redirect.
type Redirect struct{}

// Decode decodes a Redirect from JSON
func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}

	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}

	return redirect, nil
}

// Encode encodes a Redirect to JSON.
func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}

	return rawMsg, nil
}
