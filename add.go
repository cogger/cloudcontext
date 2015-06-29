package cloudcontext

import (
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/cloud"
	"gopkg.in/cogger/cloudcontext.v1/client"
)

//Add adds the cloud context to the context
func Add(projectID string) func(context.Context, *http.Request) context.Context {
	return func(ctx context.Context, req *http.Request) context.Context {
		hc := client.Get(ctx)
		return cloud.WithContext(ctx, projectID, hc)
	}
}
