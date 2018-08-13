package middleware

import (
	"net/http"

	"github.com/rs/xid"
)

const (
	// Key is the name of the response header key where we write the problem ID
	Key = "X-Problem-ID"
)

type handler struct {
	NextHandler http.Handler
}

// NewHandler is a constructor for the handler struct
func NewHandler(next http.Handler) http.Handler {
	return &handler{
		NextHandler: next,
	}
}

// ServeHTTP adds a header with a transaction ID before calling the next handler
func (middleware *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		problemID string
	)
	if len(r.Header.Get(Key)) > 0 {
		problemID = r.Header.Get(Key)
	} else {
		problemID = xid.New().String()
	}

	r.Header.Set(Key, problemID)

	middleware.NextHandler.ServeHTTP(w, r)
}
