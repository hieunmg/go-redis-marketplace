package common

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// AppError describe format error message
type AppError struct {
	StatusCode int    `json:"status_code" example:"400"`
	RootErr    error  `json:"-"`
	Message    string `json:"message"`
	Log        string `json:"log"`
	Key        string `json:"error_key"`
} // @name AppError

var (
	// RecordNotFound is error message when gorm.ErrRecordNotFound called
	RecordNotFound = errors.New("record not found")
)

// NewErrorResponse response status code 400
func NewErrorResponse(root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

// NewErrorResponse fully response with status code
func NewFullErrorResponse(statusCode int, root error, msg, log, key string) *AppError {
	return &AppError{
		StatusCode: statusCode,
		RootErr:    root,
		Message:    msg,
		Log:        log,
		Key:        key,
	}
}

// NewUnauthorized response status code 401
func NewUnauthorized(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusUnauthorized,
		RootErr:    root,
		Message:    msg,
		Log:        msg,
		Key:        key,
	}
}

// NewForbidden response status code 403
func NewForbidden(root error, msg, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusForbidden,
		RootErr:    root,
		Message:    msg,
		Log:        msg,
		Key:        key,
	}
}

// NewCustomError custom error
func NewCustomError(root error, msg string, key string) *AppError {
	if root != nil {
		return NewErrorResponse(root, msg, root.Error(), key)
	}

	return NewErrorResponse(errors.New(msg), msg, msg, key)
}

// RootError find root error
func (e *AppError) RootError() error {
	if err, ok := e.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return e.RootErr
}

// Error override error interface of go
func (e *AppError) Error() string {
	return e.RootError().Error()
}

// ErrDB defined response when calling database
func ErrDB(err error) *AppError {
	return NewErrorResponse(err, "something went wrong with DB", err.Error(), "DB_ERROR")
}

// ErrInvalidRequest invalid request
func ErrInvalidQuery(err error) *AppError {
	return NewErrorResponse(err, "invalid query", err.Error(), "ErrInvalidQuery")
}

// ErrInvalidRequest invalid request
func ErrInvalidRequest(err error) *AppError {
	return NewErrorResponse(err, "invalid request", err.Error(), "ErrInvalidRequest")
}

// ErrInternal response status code 500
func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err,
		"something went wrong in the server", err.Error(), "ErrInternal")
}

// ErrCannotListEntity list entity is error
func ErrCannotListEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot list %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotList%s", entity),
	)
}

// ErrCannotDeleteEntity delete entity is error
func ErrCannotDeleteEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot delete %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotDelete%s", entity),
	)
}

// ErrCannotDeleteEntity update entity is error
func ErrCannotUpdateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot update %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotUpdate%s", entity),
	)
}

// ErrCannotGetEntity get entity is error
func ErrCannotGetEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot get %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotGet%s", entity),
	)
}

// ErrEntityDeleted entity is deleted
func ErrEntityDeleted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s deleted", strings.ToLower(entity)),
		fmt.Sprintf("Err%sDeleted", entity),
	)
}

// ErrEntityExisted entity is existed
func ErrEntityExisted(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s already exists", strings.ToLower(entity)),
		fmt.Sprintf("Err%sAlreadyExists", entity),
	)
}

// ErrEntityNotFound entity not found
func ErrEntityNotFound(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("%s not found", strings.ToLower(entity)),
		fmt.Sprintf("Err%sNotFound", entity),
	)
}

// ErrEntityNotFound entity can not created
func ErrCannotCreateEntity(entity string, err error) *AppError {
	return NewCustomError(
		err,
		fmt.Sprintf("Cannot Create %s", strings.ToLower(entity)),
		fmt.Sprintf("ErrCannotCreate%s", entity),
	)
}

// ErrEntityNotFound no permission
func ErrNoPermission(err error) *AppError {
	return NewCustomError(
		err,
		("You have no permission"),
		("ErrNoPermission"),
	)
}

func ErrUnAuthorized(err error) *AppError {
	return NewCustomError(
		err,
		("You have no authorization"),
		("ErrUnAuthorized"),
	)
}

func ErrWrongCurrentPassword(err error) *AppError {
	return NewCustomError(
		err,
		("Current password is incorrect. Please try again"),
		("ErrWrongPassword"),
	)
}

func ErrWrongConfirmPassword(err error) *AppError {
	return NewCustomError(
		err,
		("Confirm password is incorrect. Please try again"),
		("ErrWrongConfirmPassword"),
	)
}
