package client

import (
	"net/http"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/appengine/urlfetch"
)

//Appengine adds an appengine authenicated client to the context
func Appengine(scopes ...string) func(context.Context, *http.Request) context.Context {
	return func(ctx context.Context, req *http.Request) context.Context {
		client := &http.Client{
			Transport: &oauth2.Transport{
				Source: google.AppEngineTokenSource(gaeContext, scopes...),
				Base: &urlfetch.Transport{
					Context:  gaeContext,
					Deadline: 30 * time.Second,
				},
			},
		}
		return context.WithValue(ctx, &http.Client{}, client)
	}
}
