package http

import (
	"context"
	"fmt"
	"net/http"

	api "github.com/preegnees/simpauth/internal/controller/http/api"
)

type authService struct{}
type securityServive struct{}

// HandleBearerAuth implements api.SecurityHandler
func (*securityServive) HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error) {
	fmt.Println(operationName)
	fmt.Println(t.GetToken())
	return ctx, nil
}

// TestGet implements api.Handler
func (*authService) TestGet(ctx context.Context) (api.TestGetRes, error) {
	return &api.Pet{
		Name: api.OptString{
			Value: "radmir",
			Set: true,
		},
	}, nil
}

func (*authService) MeGet(ctx context.Context) error {
	return nil
}

func Run() error {

	service := authService{}
	secService := securityServive{}

	srv, err := api.NewServer(&service, &secService)
	if err != nil {
		return err
	}
	if err := http.ListenAndServe("localhost:8080", srv); err != nil {
		return err
	}
	return nil
}

var _ api.Handler = (*authService)(nil)
var _ api.SecurityHandler = (*securityServive)(nil)
