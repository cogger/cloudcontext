// +build appengine appenginevm

package client

import (
	"net/http"

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
				Source: google.AppEngineTokenSource(ctx, scopes...),
				Base: &urlfetch.Transport{Context:  ctx},
			},
		}

		return context.WithValue(ctx, clientKey{}, client)
	}
}
