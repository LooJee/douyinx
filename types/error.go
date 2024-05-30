package types

import "errors"

var (
	ErrInvalidContentFormat = errors.New("invalid content format")
	ErrNeedAuth             = errors.New("need auth")
)
