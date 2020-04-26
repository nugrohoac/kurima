package kurima

import "errors"

// ErrNotFound .
var ErrNotFound = errors.New("not found")

// ErrorWrongPassword .
var ErrorWrongPassword = errors.New("wrong password")

// ErrInValidUser .
var ErrInValidUser = errors.New("user is not valid")

// ErrDuplicatedUser .
var ErrDuplicatedUser = errors.New("user is exist")

// ErrBindStruct .
var ErrBindStruct = errors.New("failed binding struct")

// ErrValidateStruct .
var ErrValidateStruct = errors.New("failed validate struct")
