package middleware

import (
	"context"
	"github.com/google/uuid"
	"net/http"
	"time"
)

type ContextKey string

const (
	RequestIDKey  ContextKey = "request_id"
	CorrelationID ContextKey = "correlation_id"
)

// Middleware Function that adds a request context to the context. setting up request id and correlation id
func WithRequestContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		requestID := uuid.New().String()
		ctx = context.WithValue(ctx, "request_id", requestID)

		correlationID := request.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = uuid.New().String()
		}
		ctx = context.WithValue(ctx, "correlation_id", correlationID)

		request = request.WithContext(ctx)

		responseWriter.Header().Set("X-Request-ID", requestID)
		responseWriter.Header().Set("X-Correlation-ID", correlationID)

		next.ServeHTTP(responseWriter, request)
	})
}

// middleware that sets a timeout for requests that can be set on a per subrouter bases
func WithTimeout(duration time.Duration) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(responseWriter http.ResponseWriter, request *http.Request) {
			ctx, cancel := context.WithTimeout(request.Context(), duration)
			defer cancel()

			request = request.WithContext(ctx)
			next.ServeHTTP(responseWriter, request)
		})
	}
}
