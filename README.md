# googlecontext [![GoDoc](https://godoc.org/github.com/cogger/googlecontext?status.png)](http://godoc.org/github.com/cogger/googlecontext)

gae addes the appengine context to the httpHandler

## Usage
~~~ go
// main.go
package main

import (
	"net/http"
	"github.com/cogger/cogger"
	"github.com/cogger/cloudcontext"
	"github.com/cogger/cloudcontext/client"
	"github.com/cogger/cloudcontext/bq"
	"golang.org/x/net/context"
)

func foo(ctx context.Context, w http.ResponseWriter, r *http.Request) int{
	service := bq.FromContext(ctx)
	//Do something with bigquery
	return http.StatusOK
}


func init() {
	fooHandler := cogger.NewHandler()
	fooHandler.AddContext(client.Compute)
	fooHandler.AddContext(cloudcontext.Add)
	fooHandler.SetHandler(foo)

  	http.Handle("/foo", fooHandler)
}

~~~
