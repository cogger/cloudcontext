package cloudcontext

import (
	"errors"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/cloud"
)

//ErrNoHTTPClient denotes that not http.Client has been added to the context
var ErrNoHTTPClient = errors.New("no http context")

//Add adds the cloud context to the context
func Add(projectID string) func(context.Context, *http.Request) context.Context {
	return func(ctx context.Context, req *http.Request) context.Context {
		client, ok := ctx.Value(&http.Client{}).(*http.Client)
		if !ok {
			panic(ErrNoHTTPClient)
		}

		return cloud.WithContext(ctx, projectID, client)
	}
}
