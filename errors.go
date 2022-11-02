package fluor

import (
	"context"
	"time"
)

type FluentError struct {
	kind           Kind
	timestamp      time.Time
	httpStatusCode int
	message        string
	wrappedError   error
	fields         H
}

type H map[string]any

type stdError interface {
	error
	Unwrap() error
	Is(error) bool
}

var _ stdError = &FluentError{}

// New initializes a new error
// Example usage:
//
//	    return fluor.New().
//		  WithError(err).
//		  WithHTTPStatusCode(http.StatusBadRequest).
//		  WithMessage("Invalid payload").
//		  Log(ctx, crewerr.LevelError)
func New() *FluentError {
	return &FluentError{
		kind:      KindDefault,
		timestamp: time.Now(),
		fields:    H{},
	}
}

func (err *FluentError) WithKind(kind Kind) *FluentError {
	err.kind = kind
	return err
}

func (err *FluentError) WithHTTPStatusCode(httpStatusCode int) *FluentError {
	err.httpStatusCode = httpStatusCode
	return err
}

func (err *FluentError) WithMessage(message string) *FluentError {
	err.message = message
	return err
}

func (err *FluentError) WithError(wrappedError error) *FluentError {
	if wrappedFluentError, ok := wrappedError.(*FluentError); ok {
		*err = *wrappedFluentError
	} else {
		err.wrappedError = wrappedError
	}
	return err
}

func (err *FluentError) WithParsedError(wrappedError error) *FluentError {
	return Parser(wrappedError, err).WithError(wrappedError)
}

func (err *FluentError) WithFields(fields H) *FluentError {
	if err.fields == nil {
		err.fields = H{}
	}
	for k, v := range fields {
		err.fields[k] = v
	}
	return err
}

func (err *FluentError) Error() string {
	s := err.Message()
	if err.wrappedError != nil {
		s += ": " + err.wrappedError.Error()
	}
	return s
}

func (err *FluentError) Kind() Kind {
	if err.kind == nil {
		return KindDefault
	}
	return err.kind
}

func (err *FluentError) Timestamp() time.Time {
	if err.timestamp.IsZero() {
		return time.Now()
	}
	return err.timestamp
}

func (err *FluentError) Log(ctx context.Context, level LogLevel) *FluentError {
	return Logger(ctx, level, err)
}

func (err *FluentError) HTTPStatusCode() int {
	if err.httpStatusCode == 0 {
		return defaultKindToHTTPStatusCode(err.Kind())
	}
	return err.httpStatusCode
}

func (err *FluentError) Message() string {
	if err.message == "" {
		return "server error"
	}
	return err.message
}

func (err *FluentError) Fields() H {
	if err.fields == nil {
		return H{}
	}
	return err.fields
}

// Compatibility with [errors.Is], [errors.As] and [errors.Unwrap]

func (err *FluentError) Is(target error) bool {
	return err.kind == target
}

func (err *FluentError) Unwrap() error {
	return err.wrappedError
}
