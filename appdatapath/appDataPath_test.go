package appdatapath

import (
	"github.com/golang-interfaces/vos"
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
			Expect(new(vos.Fake)).Should(Not(BeNil()))
		})
	})
})
