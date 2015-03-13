package client

import (
	"net/http"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

//Compute adds the compute authenticated client to the context
func Compute(ctx context.Context, req *http.Request) context.Context {
	client := &http.Client{
		Transport: &oauth2.Transport{
			Source: google.ComputeTokenSource(""),
		},
	}
	return context.WithValue(ctx, clientKey{}, client)
}
