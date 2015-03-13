package client

import (
	"errors"

	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/appengine/log"
)

//ErrNoHTTPClient denotes that not http.Client has been added to the context
var ErrNoHTTPClient = errors.New("no http context")

type clientKey struct{}

//Get gets the http.Client from the ctx
func Get(ctx context.Context) *http.Client {
	log.Infof(ctx, "%+v", ctx.Value(http.Client{}))
	client, ok := ctx.Value(clientKey{}).(*http.Client)
	if !ok {
		panic(ErrNoHTTPClient)
	}
	return client
}
