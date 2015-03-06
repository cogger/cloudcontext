# cloudcontext [![GoDoc](https://godoc.org/github.com/cogger/cloudcontext?status.png)](http://godoc.org/github.com/cogger/cloudcontext)

cloudcontext adds the cloud context to the context.  There are also functions for adding authenticated http.Client to the context.  

## Usage
~~~ go
// main.go
package main

import (
	"log"
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
	fooHandler.AddContext(client.Compute, cloudcontext.Add)
	fooHandler.SetHandler(foo)

  	http.Handle("/foo", fooHandler)
  	log.Fatal(http.ListenAndServe(":8080", nil))
}

~~~
