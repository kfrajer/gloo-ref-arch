package petclinic_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/solo-io/gloo-ref-arch/petclinic"
	"github.com/solo-io/go-utils/testutils"
	"testing"
)

func TestPetclinic(t *testing.T) {
	RegisterFailHandler(Fail)
	testutils.RegisterPreFailHandler(
		func() {
			testutils.PrintTrimmedStack()
		})
	testutils.RegisterCommonFailHandlers()
	RunSpecs(t, "Petclinic Suite")
}

var _ = Describe("petclinic", func() {
	testWorkflow := petclinic.GetTestWorkflow()

	BeforeSuite(func() {
		testWorkflow.Setup(".")
	})

	It("runs", func() {
		testWorkflow.Run(".")
	})
})
