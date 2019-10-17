package main

import (
	"net/http"

	"github.com/go-kit/kit/endpoint"
)

// Middleware export
type Middleware func(endpoint.Endpoint) endpoint.Endpoint

// AuthMiddlewareType export
type AuthMiddlewareType func(http.Handler) http.Handler
