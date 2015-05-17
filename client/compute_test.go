package client_test

import (
	"net/http"

	. "github.com/cogger/cloudcontext/client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
)

var _ = Describe("Compute", func() {
	It("should return a client", func() {
		req, err := http.NewRequest("GET", "/", nil)
		Expect(err).ToNot(HaveOccurred())
		ctx := Compute(context.Background(), req)
		Expect(Get(ctx)).ToNot(BeNil())
	})
})
