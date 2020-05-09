package kurima

import (
	"fmt"
)

// ErrNotFound .
type ErrNotFound struct {
	Message string
}

func (enf ErrNotFound) Error() string {
	return enf.Message
}

// ErrInValid .
type ErrInValid struct {
	Message string
}

func (ei ErrInValid) Error() string {
	return ei.Message
}

// ErrDuplicated .
type ErrDuplicated struct {
	Message string
}

func (ed ErrDuplicated) Error() string {
	return ed.Message
}

// ErrBindStruct .
type ErrBindStruct struct {
	Message string
}

func (ebs ErrBindStruct) Error() string {
	return ebs.Message
}

// ErrValidateStruct .
type ErrValidateStruct struct {
	Message string
}

func (evs ErrValidateStruct) Error() string {
	return evs.Message
}

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

// ErrNoRowAffected .
type ErrNoRowAffected struct {
	Message string
}

func (enr ErrNoRowAffected) Error() string {
	return enr.Message
}
