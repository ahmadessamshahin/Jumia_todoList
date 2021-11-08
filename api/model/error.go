package model

import "errors"

var ErrInvalidInput = errors.New("INVALID INPUT")

var ErrInvalidQuery = func(msg string) { errors.New(msg) }
