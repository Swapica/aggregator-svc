package middleware

import (
	"context"
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func EndpointBuilder(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		route := chi.URLParam(r, "*")
		ctx := context.WithValue(r.Context(), "actionTarget", fmt.Sprintf("v1/%v", route))

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
