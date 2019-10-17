package main

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
)

// AuthMiddleware auth
func AuthMiddleware(logger log.Logger) Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, r interface{}) (interface{}, error) {
			logger.Log("Request:", r)
			// e := MakeUnauthenticatedEndPoint
			var e endpoint.Endpoint
			e = MakeUnauthenticatedEndPoint()
			return e(ctx, r)
		}
	}
}

// TestAuthMiddleware export
func testAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}

// TestAuthMiddleware export
func TestAuthMiddleware(logger log.Logger) AuthMiddlewareType {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Log("Auth", "Testing authentication")
			logger.Log("Header", r.Header)
			next.ServeHTTP(w, r)
		})
	}
}
