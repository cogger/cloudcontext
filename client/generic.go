package client

import (
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
)

//Generic adds a generic oauth authenticated client to the context
func Generic(config *oauth2.Config) func(context.Context, *http.Request) context.Context {
	return func(ctx context.Context, req *http.Request) context.Context {
		return context.WithValue(ctx, &http.Client{}, config.Client(ctx, nil))
	}
}
