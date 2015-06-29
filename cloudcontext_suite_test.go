package cloudcontext

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestCloudcontext(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Cloudcontext Suite")
}
