package bq

import (
	"errors"
	"net/http"

	"golang.org/x/net/context"
	"google.golang.org/api/bigquery/v2"
	"google.golang.org/cloud/internal"
)

//ErrNoBigQuery denotes that the big query context is not added correctly
var ErrNoBigQuery = errors.New("No Big Query")

//FromContext gets the big query service from the context
func FromContext(ctx context.Context) *bigquery.Service {
	service, ok := internal.Service(ctx, "bigquery", func(hc *http.Client) interface{} {
		svc, err := bigquery.New(hc)
		if err != nil {
			panic(ErrNoBigQuery)
		}
		return svc
	}).(*bigquery.Service)

	if !ok {
		panic(ErrNoBigQuery)
	}
	return service
}

func FromContextSafe(ctx context.Context) (*bigquery.Service, error) {
	service, ok := internal.Service(ctx, "bigquery", func(hc *http.Client) interface{} {
		svc, _ := bigquery.New(hc)
		return svc
	}).(*bigquery.Service)

	if !ok {
		return nil, ErrNoBigQuery
	}
	return service, nil
}
