package client_test

import (
	"net/http"

	. "github.com/cogger/cloudcontext/client"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Generic", func() {
	It("should return a client with ", func() {
		req, err := http.NewRequest("GET", "/", nil)
		Expect(err).ToNot(HaveOccurred())
		ctx := Generic(&oauth2.Config{})(context.Background(), req)
		Expect(Get(ctx)).ToNot(BeNil())
	})
})
