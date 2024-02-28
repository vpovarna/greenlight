package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var ErrInvalidFormat = errors.New("invalid runtime format")

type Runtime int32

func (r Runtime) MarshalJSON() ([]byte, error) {
	jsonValue := fmt.Sprintf("%d mins", r)
	quotedJsonValue := strconv.Quote(jsonValue)
	return []byte(quotedJsonValue), nil
}

func (r *Runtime) UnmarshalJSON(jsonValue []byte) error {

	unquotedJson, err := strconv.Unquote(string(jsonValue))
	if err != nil {
		return ErrInvalidFormat
	}

	parts := strings.Split(unquotedJson, " ")

	if len(parts) != 2 || parts[1] != "mins" {
		return ErrInvalidFormat
	}

	i, err := strconv.ParseInt(parts[0], 10, 32)
	if err != nil {
		return ErrInvalidFormat
	}

	// Convert the int32 to a Runtime type and assign this to the receiver. Note that we
	// use the * operator to deference the receiver (which is a pointer to a Runtime
	// type) in order to set the underlying value of the pointer.
	*r = Runtime(i)

	return nil
}
