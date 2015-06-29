package bq

import (
	"log"
	"net/http"

	"gopkg.in/cogger/cloudcontext.v1/client"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"gopkg.in/cogger/cloudcontext.v1"
)

var _ = Describe("Bq", func() {
	It("blah", func() {
		req, err := http.NewRequest("GET", "/", nil)
		Expect(err).ToNot(HaveOccurred())
		ctx := client.Generic(&oauth2.Config{})(context.Background(), req)
		service := FromContext(cloudcontext.Add("someid")(ctx, req))
		log.Println(service)
	})
})
