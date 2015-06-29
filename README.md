# cloudcontext 

[![GoDoc](https://godoc.org/gopkg.in/cogger/cloudcontext.v1?status.png)](http://godoc.org/gopkg.in/cogger/cloudcontext.v1)  
[![Build Status](https://travis-ci.org/cogger/cloudcontext.svg?branch=master)](https://travis-ci.org/cogger/cloudcontext)  
[![Coverage Status](https://coveralls.io/repos/cogger/cloudcontext/badge.svg?branch=master)](https://coveralls.io/r/cogger/cloudcontext?branch=master)  
[![License](http://img.shields.io/:license-apache-blue.svg)](http://www.apache.org/licenses/LICENSE-2.0.html)


cloudcontext adds the cloud context to the context.  There are also functions for adding authenticated http.Client to the context.  

## Installation

The import path for the package is *gopkg.in/cogger/cloudcontext.v1*.

To install it, run:

    go get gopkg.in/cogger/cloudcontext.v1

## Usage
~~~ go
// main.go
package main

import (
	"log"
	"net/http"
	"gopkg.in/cogger/cogger.v1"
	"gopkg.in/cogger/cloudcontext.v1"
	"gopkg.in/cogger/cloudcontext.v1/client"
	"gopkg.in/cogger/cloudcontext.v1/bq"
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
