package path

import (
	"github.com/appdataspec/sdk-golang/util/vos"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("appdata", func() {
	Context("New", func() {
		It("should return AppData", func() {
			/* arrange/act/assert */
			Expect(New()).Should(Not(BeNil()))
		})
	})
	Context("_new", func() {
		It("should return AppData", func() {
			/* arrange/act/assert */
			Expect(new(vos.FakeVos)).Should(Not(BeNil()))
		})
	})
})
