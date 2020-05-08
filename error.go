package kurima

import (
	"errors"
	"fmt"
)

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

// InternalError .
type InternalError struct {
	Path string
}

func (i InternalError) Error() string {
	return fmt.Sprintf("internal server error at : %v", i.Path)
}

// SyntaxError .
type SyntaxError struct {
	Line int
	Col  int
}

func (s SyntaxError) Error() string {
	return fmt.Sprintf("%d:%d: syntax error", s.Line, s.Col)
}

// ErrorAuth is used to return error auth
type ErrorAuth struct {
	Message string
}

func (ea ErrorAuth) Error() string {
	return ea.Message
}
