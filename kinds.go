package fluor

import (
	"errors"
	"net/http"
)

type Kind error

var (
	KindDefault  Kind = errors.New("default")
	KindNotFound Kind = errors.New("not found")
	// TODO: add more kinds
)

func defaultKindToHTTPStatusCode(kind Kind) int {
	switch kind {
	case KindNotFound:
		return http.StatusNotFound

	default:
		return http.StatusInternalServerError
	}
}
