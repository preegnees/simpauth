// Code generated by ogen, DO NOT EDIT.

package api

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// MeGet implements GET /me operation.
//
// GET /me
func (UnimplementedHandler) MeGet(ctx context.Context) error {
	return ht.ErrNotImplemented
}

// TestGet implements GET /test operation.
//
// GET /test
func (UnimplementedHandler) TestGet(ctx context.Context) (r TestGetRes, _ error) {
	return r, ht.ErrNotImplemented
}
