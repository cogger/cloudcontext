package bq

import (
	"net/http"

	"gopkg.in/cogger/cloudcontext.v1/client"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"gopkg.in/cogger/cloudcontext.v1"
)

var _ = Describe("Bq", func() {
	It("should return the context when it exists", func() {
		req, err := http.NewRequest("GET", "/", nil)
		Expect(err).ToNot(HaveOccurred())

		ctx := client.Generic(&oauth2.Config{})(context.Background(), req)

		Expect(func() {
			FromContext(cloudcontext.Add("someid")(ctx, req))
		}).NotTo(Panic())
	})

	It("should panic when there is no client", func() {
		req, err := http.NewRequest("GET", "/", nil)
		Expect(err).ToNot(HaveOccurred())

		ctx := context.Background()

		Expect(func() {
			FromContext(cloudcontext.Add("someid")(ctx, req))
		}).To(Panic())
	})

	It("should not return a serivce and nil when it exists", func() {
		req, err := http.NewRequest("GET", "/", nil)
		Expect(err).ToNot(HaveOccurred())

		ctx := client.Generic(&oauth2.Config{})(context.Background(), req)

		service, err := FromContextSafe(cloudcontext.Add("someid")(ctx, req))
		Expect(service).ToNot(BeNil())
		Expect(err).To(BeNil())
	})
})
