package api

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
)

// This exists to copy the body from the request into a context variable
// This is needed because the csrf middleware will consume the body, so we need a persisted version in context
func CopyBody(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		buf, _ := ioutil.ReadAll(r.Body)

		rdr := ioutil.NopCloser(bytes.NewBuffer(buf))
		ctx := context.WithValue(r.Context(), "rBody", buf)
		r.Body = rdr
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}