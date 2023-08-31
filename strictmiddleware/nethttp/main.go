package nethttp

import (
	"context"
	"net/http"
)

type StrictHTTPHandlerFunc func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (response interface{}, err error)

type StrictHTTPMiddlewareFunc func(f StrictHTTPHandlerFunc, operationID string) StrictHTTPHandlerFunc
