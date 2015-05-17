package bq_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBq(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bq Suite")
}
