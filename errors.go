package wire

import (
	"errors"
)

var (
	ErrRequestHeader  = errors.New("wire: malformed request header")
	ErrRequestVersion = errors.New("wire: invalid or unsupported protocol version in request header")

	ErrResponseHeader  = errors.New("wire: malformed response header")
	ErrResponseVersion = errors.New("wire: invalid or unsupported protocol version in response header")

	ErrInvalidChunkedEncoding = errors.New("wire: invalid chunked encoding")
)

var (
	errMalformedHeader = errors.New("malformed header")
	errInvalidVersion  = errors.New("invalid version")
)
