package json

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/svelama/microsvc/shortener"
)

type Redirect struct{}

// r redirect
func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {

	rawMsg, err := json.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}

	return rawMsg, nil
}

func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error){

	msgObject := &shortener.Redirect{}
	if err := json.Unmarshal(input, msgObject); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return msgObject, nil
}
