package rest

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestTransformation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Transformation Suite")
}