package cloudcontext

import (
	"net/http"

	"github.com/cogger/cloudcontext/client"
	"golang.org/x/net/context"
	"google.golang.org/cloud"
)

//Add adds the cloud context to the context
func Add(projectID string) func(context.Context, *http.Request) context.Context {
	return func(ctx context.Context, req *http.Request) context.Context {
		return cloud.WithContext(ctx, projectID, client.Get(ctx))
	}
}
