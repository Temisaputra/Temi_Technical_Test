package helper

import "errors"

// Error constants (for comparison via errors.Is)
var (
	ErrTokenExpired            = errors.New("token expired")
	ErrInsufficientPermissions = errors.New("insufficient permissions")
	ErrUnauthorizedBase        = errors.New("unauthorized")
	ErrBadRequestBase          = errors.New("bad request")
	ErrNotFoundBase            = errors.New("not found")
	ErrForbiddenBase           = errors.New("forbidden")
)

// ---------- Bad Request ----------
type ErrBadRequest struct {
	Message string
	cause   error
}

func NewErrBadRequest(message string) *ErrBadRequest {
	return &ErrBadRequest{
		Message: message,
		cause:   ErrBadRequestBase,
	}
}

func (e *ErrBadRequest) Error() string {
	return e.Message
}

func (e *ErrBadRequest) Unwrap() error {
	return e.cause
}

// ---------- Not Found ----------
type ErrNotFound struct {
	Message string
	cause   error
}

func NewErrNotFound(message string) *ErrNotFound {
	return &ErrNotFound{
		Message: message,
		cause:   ErrNotFoundBase,
	}
}

func (e *ErrNotFound) Error() string {
	return e.Message
}

func (e *ErrNotFound) Unwrap() error {
	return e.cause
}

// ---------- Unauthorized ----------
type ErrUnauthorized struct {
	Message string
	cause   error
}

func NewErrUnauthorized(message string) *ErrUnauthorized {
	var cause error
	switch message {
	case "token expired":
		cause = ErrTokenExpired
	default:
		cause = ErrUnauthorizedBase
	}

	return &ErrUnauthorized{
		Message: message,
		cause:   cause,
	}
}

func (e *ErrUnauthorized) Error() string {
	return e.Message
}

func (e *ErrUnauthorized) Unwrap() error {
	return e.cause
}

// ---------- Forbidden ----------
type ErrForbidden struct {
	Message string
	cause   error
}

func NewErrForbidden(message string) *ErrForbidden {
	var cause error
	switch message {
	case "insufficient permissions":
		cause = ErrInsufficientPermissions
	default:
		cause = ErrForbiddenBase
	}

	return &ErrForbidden{
		Message: message,
		cause:   cause,
	}
}

func (e *ErrForbidden) Error() string {
	return e.Message
}

func (e *ErrForbidden) Unwrap() error {
	return e.cause
}
