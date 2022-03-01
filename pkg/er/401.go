package er

import "net/http"

var (
	// ErrMissingToken means Must provide Authorization header with format `Bearer {token}`
	ErrMissingToken = NewAPPError(http.StatusUnauthorized, 40100, "Must provide Authorization header with format `Bearer {token}`")

	// ErrAuthHeaderFormat means must provide Authorization header with format `Bearer {token}`
	ErrAuthHeaderFormat = NewAPPError(http.StatusUnauthorized, 40101, "Must provide Authorization header with format `Bearer {token}`")
)
