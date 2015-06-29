package cloudcontext

import (
	"reflect"

	"net/http"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"gopkg.in/cogger/cloudcontext.v1/client"
)

var _ = Describe("Add", func() {
	contextInterface := reflect.TypeOf((*context.Context)(nil)).Elem()

	It("should return a client", func() {
		req, err := http.NewRequest("GET", "/", nil)
		Expect(err).ToNot(HaveOccurred())
		ctx := client.Generic(&oauth2.Config{})(context.Background(), req)
		Expect(client.Get(ctx)).ToNot(BeNil())
		Expect(reflect.TypeOf(Add("someid")(ctx, req)).Implements(contextInterface)).To(BeTrue())
	})
})
